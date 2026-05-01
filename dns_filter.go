package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

type DNSConfig struct {
	Enabled           bool   `json:"enabled"`
	Workers           int    `json:"workers"`
	TimeoutSeconds    int    `json:"timeout_seconds"`
	CacheFile         string `json:"cache_file"`
	DeadDomainsFile   string `json:"dead_domains_file"`
	RecheckAfterHours int    `json:"recheck_after_hours"`
	UsePuredns        bool   `json:"use_puredns"`
	PurednsResolvers  string `json:"puredns_resolvers"`
}

type CacheEntry struct {
	Status    string `json:"status"`
	CheckedAt string `json:"checked_at"`
}

func LoadDNSConfig(path string) DNSConfig {
	defaultConfig := DNSConfig{
		Enabled:           true,
		Workers:           100,
		TimeoutSeconds:    3,
		CacheFile:         "rules/dns_cache.json",
		DeadDomainsFile:   "rules/dead_domains.txt",
		RecheckAfterHours: 24,
		UsePuredns:        true,
		PurednsResolvers:  "resolvers.txt",
	}

	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("DNS config not found, using defaults:", err)
		return defaultConfig
	}

	err = json.Unmarshal(data, &defaultConfig)
	if err != nil {
		fmt.Println("Error parsing DNS config, using defaults:", err)
		return defaultConfig
	}

	return defaultConfig
}

func LoadDomainCache(path string) map[string]CacheEntry {
	cache := make(map[string]CacheEntry)

	data, err := os.ReadFile(path)
	if err != nil {
		return cache
	}

	err = json.Unmarshal(data, &cache)
	if err != nil {
		fmt.Println("Error parsing cache file, starting fresh:", err)
		return cache
	}

	return cache
}

func SaveDomainCache(cache map[string]CacheEntry, path string) error {
	dir := filepath.Dir(path)
	if dir != "" && dir != "." {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return fmt.Errorf("failed to create cache directory: %v", err)
		}
	}

	data, err := json.MarshalIndent(cache, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal cache: %v", err)
	}

	err = os.WriteFile(path, data, 0644)
	if err != nil {
		return fmt.Errorf("failed to write cache file: %v", err)
	}

	return nil
}

var domainPatternRegex = regexp.MustCompile(`^\|\|([a-zA-Z0-9\-_.]+)\^`)

func ExtractDomains(rules []string) []string {
	seen := make(map[string]bool)
	var domains []string

	for _, rule := range rules {
		rule = strings.TrimSpace(rule)
		if rule == "" || strings.HasPrefix(rule, "#") || strings.HasPrefix(rule, "!") || strings.HasPrefix(rule, "[") {
			continue
		}

		matches := domainPatternRegex.FindStringSubmatch(rule)
		if len(matches) > 1 {
			domain := strings.ToLower(strings.TrimSpace(matches[1]))
			if domain != "" && !seen[domain] {
				seen[domain] = true
				domains = append(domains, domain)
			}
		}
	}

	return domains
}

func IsStale(entry CacheEntry, recheckAfterHours int) bool {
	if entry.CheckedAt == "" {
		return true
	}
	checkedAt, err := time.Parse(time.RFC3339, entry.CheckedAt)
	if err != nil {
		return true
	}
	return time.Since(checkedAt).Hours() > float64(recheckAfterHours)
}

func IsPurednsInstalled() bool {
	_, err := exec.LookPath("puredns")
	return err == nil
}

func InstallPuredns() error {
	if IsPurednsInstalled() {
		fmt.Println("puredns is already installed")
		return nil
	}

	fmt.Println("Installing puredns...")

	cmd := exec.Command("go", "install", "github.com/d3mondev/puredns/v2@latest")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to install puredns: %v", err)
	}

	// Verify installation
	if !IsPurednsInstalled() {
		return fmt.Errorf("puredns installed but not found in PATH")
	}

	fmt.Println("puredns installed successfully")
	return nil
}

func CreateResolversFile(path string) error {
	resolvers := []string{
		"8.8.8.8",
		"8.8.4.4",
		"1.1.1.1",
		"1.0.0.1",
		"9.9.9.9",
		"149.112.112.112",
		"208.67.222.222",
		"208.67.220.220",
	}

	content := strings.Join(resolvers, "\n") + "\n"
	return os.WriteFile(path, []byte(content), 0644)
}

func RunPuredns(domains []string, config DNSConfig) ([]string, error) {
	fmt.Printf("Running puredns with %d domains...\n", len(domains))

	tmpDir, err := os.MkdirTemp("", "puredns_*")
	if err != nil {
		return nil, fmt.Errorf("failed to create temp directory: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	domainsFile := filepath.Join(tmpDir, "domains.txt")
	err = os.WriteFile(domainsFile, []byte(strings.Join(domains, "\n")), 0644)
	if err != nil {
		return nil, fmt.Errorf("failed to write domains file: %v", err)
	}

	resolversFile := config.PurednsResolvers
	if _, err := os.Stat(resolversFile); os.IsNotExist(err) {
		err = CreateResolversFile(resolversFile)
		if err != nil {
			return nil, fmt.Errorf("failed to create resolvers file: %v", err)
		}
	}

	aliveFile := filepath.Join(tmpDir, "alive.txt")
	deadFile := filepath.Join(tmpDir, "dead.txt")

	cmd := exec.Command("puredns", "resolve", domainsFile,
		"--resolvers", resolversFile,
		"--quiet",
		"--write", aliveFile,
		"--write-failures", deadFile,
	)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	start := time.Now()
	err = cmd.Run()
	elapsed := time.Since(start)

	if err != nil {
		fmt.Printf("puredns completed with exit code: %v (took %s)\n", err, elapsed.Round(time.Second))
	} else {
		fmt.Printf("puredns completed successfully (took %s)\n", elapsed.Round(time.Second))
	}

	deadDomains := readLinesFromFile(deadFile)
	aliveDomains := readLinesFromFile(aliveFile)

	fmt.Printf("puredns results: %d alive, %d dead\n", len(aliveDomains), len(deadDomains))

	return deadDomains, nil
}

func readLinesFromFile(path string) []string {
	var lines []string

	data, err := os.ReadFile(path)
	if err != nil {
		return lines
	}

	scanner := bufio.NewScanner(strings.NewReader(string(data)))
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" {
			lines = append(lines, line)
		}
	}

	return lines
}

func RunDNSFilter(inputFile string, outputFile string, configPath string) {
	config := LoadDNSConfig(configPath)

	if !config.Enabled {
		fmt.Println("DNS filtering is disabled in config")
		data, err := os.ReadFile(inputFile)
		if err != nil {
			fmt.Println("Error reading input:", err)
			return
		}
		if err := os.WriteFile(outputFile, data, 0644); err != nil {
			fmt.Println("Error writing output:", err)
		}
		return
	}

	fmt.Println("Loading rules from:", inputFile)
	data, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	lines := strings.Split(string(data), "\n")
	fmt.Printf("Total rules: %d\n", len(lines))

	domains := ExtractDomains(lines)
	fmt.Printf("Unique domains extracted: %d\n", len(domains))

	if len(domains) == 0 {
		fmt.Println("No domains found to check, skipping DNS filter")
		if err := os.WriteFile(outputFile, data, 0644); err != nil {
			fmt.Println("Error writing output:", err)
		}
		return
	}

	cache := LoadDomainCache(config.CacheFile)
	fmt.Printf("Cache entries loaded: %d\n", len(cache))

	var toCheck []string
	var deadFromCache []string
	var aliveFromCache []string

	for _, domain := range domains {
		entry, ok := cache[domain]
		if ok && !IsStale(entry, config.RecheckAfterHours) {
			if entry.Status == "dead" {
				deadFromCache = append(deadFromCache, domain)
			} else {
				aliveFromCache = append(aliveFromCache, domain)
			}
		} else {
			toCheck = append(toCheck, domain)
		}
	}

	fmt.Printf("Cached (alive): %d\n", len(aliveFromCache))
	fmt.Printf("Cached (dead): %d\n", len(deadFromCache))
	fmt.Printf("Need to check: %d\n", len(toCheck))

	var newDeadDomains []string

	if len(toCheck) > 0 {
		if config.UsePuredns {
			if err := InstallPuredns(); err != nil {
				fmt.Printf("Failed to install puredns: %v\n", err)
				fmt.Println("Skipping DNS filtering, writing original rules")
				if err := os.WriteFile(outputFile, data, 0644); err != nil {
					fmt.Println("Error writing output:", err)
				}
				return
			}

			newDeadDomains, err = RunPuredns(toCheck, config)
			if err != nil {
				fmt.Printf("puredns error: %v\n", err)
				fmt.Println("Skipping DNS filtering, writing original rules")
				if err := os.WriteFile(outputFile, data, 0644); err != nil {
					fmt.Println("Error writing output:", err)
				}
				return
			}

			deadSet := make(map[string]bool)
			for _, d := range newDeadDomains {
				deadSet[d] = true
			}

			now := time.Now().UTC().Format(time.RFC3339)
			for _, domain := range toCheck {
				status := "alive"
				if deadSet[domain] {
					status = "dead"
				}
				cache[domain] = CacheEntry{
					Status:    status,
					CheckedAt: now,
				}
			}
		} else {
			fmt.Println("puredns is disabled. For 1M+ domains, enable use_puredns in dns_config.json")
			fmt.Println("Writing original rules without DNS filtering")
			if err := os.WriteFile(outputFile, data, 0644); err != nil {
				fmt.Println("Error writing output:", err)
			}
			return
		}
	}

	allDeadDomains := append(deadFromCache, newDeadDomains...)

	fmt.Printf("Total dead domains: %d\n", len(allDeadDomains))

	if err := SaveDomainCache(cache, config.CacheFile); err != nil {
		fmt.Printf("Warning: Failed to save cache: %v\n", err)
	} else {
		fmt.Println("Cache saved to:", config.CacheFile)
	}

	deadDomainsContent := fmt.Sprintf("# Dead domains list - generated %s\n", time.Now().UTC().Format(time.RFC3339))
	deadDomainsContent += fmt.Sprintf("# Total: %d domains\n", len(allDeadDomains))
	deadDomainsContent += "# These domains returned NXDOMAIN or failed DNS resolution\n#\n"
	deadDomainsContent += strings.Join(allDeadDomains, "\n") + "\n"

	deadDir := filepath.Dir(config.DeadDomainsFile)
	if deadDir != "" && deadDir != "." {
		os.MkdirAll(deadDir, 0755)
	}
	if err := os.WriteFile(config.DeadDomainsFile, []byte(deadDomainsContent), 0644); err != nil {
		fmt.Printf("Warning: Failed to save dead domains: %v\n", err)
	} else {
		fmt.Println("Dead domains saved to:", config.DeadDomainsFile)
	}

	deadSet := make(map[string]bool)
	for _, d := range allDeadDomains {
		deadSet[d] = true
	}

	var filtered []string
	removedCount := 0
	for _, rule := range lines {
		ruleTrimmed := strings.TrimSpace(rule)
		if ruleTrimmed == "" || strings.HasPrefix(ruleTrimmed, "#") || strings.HasPrefix(ruleTrimmed, "!") || strings.HasPrefix(ruleTrimmed, "[") {
			filtered = append(filtered, rule)
			continue
		}

		matches := domainPatternRegex.FindStringSubmatch(ruleTrimmed)
		if len(matches) > 1 {
			domain := strings.ToLower(strings.TrimSpace(matches[1]))
			if deadSet[domain] {
				removedCount++
				continue
			}
		}

		filtered = append(filtered, rule)
	}

	fmt.Printf("Rules after filtering: %d (removed %d)\n", len(filtered), removedCount)

	output := strings.Join(filtered, "\n")
	if err := os.WriteFile(outputFile, []byte(output), 0644); err != nil {
		fmt.Println("Error writing output:", err)
		return
	}
	fmt.Println("Filtered rules written to:", outputFile)
}
