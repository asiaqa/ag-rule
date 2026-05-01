[![Ad-list-processing](https://github.com/asiaqa/ag-rule/actions/workflows/ad-list.yml/badge.svg)](https://github.com/asiaqa/ag-rule/actions/workflows/ad-list.yml)

# ag-rule

Consolidated adblock rules for AdGuard Home, Pi-hole, dnsmasq, and Unbound. Updated 3 times daily.

整合多個來源的廣告封鎖規則，適用於 AdGuard Home、Pi-hole、dnsmasq 及 Unbound。每日更新 3 次。

---

## Features / 功能

| Feature | Description |
|---------|-------------|
| Multi-source | Consolidates 30+ filter lists into one |
| Multi-format | AdGuard Home, hosts, dnsmasq, Unbound, uBlock Origin, domain-only |
| 4 tiers | Super, Full, Medium, Min — choose your level |
| DNS filtering | Removes dead domains (NXDOMAIN) to reduce file size |
| Transparency | Dead domains list published for review |

---

## Download Links / 下載連結

### AdGuard Home Format

| Tier | Raw URL | CDN (12h delay) |
|------|---------|-----------------|
| Super | [output_super.txt](https://raw.githubusercontent.com/asiaqa/ag-rule/main/rules/output_super.txt) | [jsdelivr](https://cdn.jsdelivr.net/gh/asiaqa/ag-rule@release/output_super.txt) |
| Full | [output_full.txt](https://raw.githubusercontent.com/asiaqa/ag-rule/main/rules/output_full.txt) | [jsdelivr](https://cdn.jsdelivr.net/gh/asiaqa/ag-rule@release/output_full.txt) |
| Medium | [output_medium.txt](https://raw.githubusercontent.com/asiaqa/ag-rule/main/rules/output_medium.txt) | [jsdelivr](https://cdn.jsdelivr.net/gh/asiaqa/ag-rule@release/output_medium.txt) |
| Min | [output_min.txt](https://raw.githubusercontent.com/asiaqa/ag-rule/main/rules/output_min.txt) | [jsdelivr](https://cdn.jsdelivr.net/gh/asiaqa/ag-rule@release/output_min.txt) |

### Hosts Format (Pi-hole)

| Tier | Raw URL | CDN |
|------|---------|-----|
| Super | [output_super_hosts.txt](https://raw.githubusercontent.com/asiaqa/ag-rule/main/rules/output_super_hosts.txt) | [jsdelivr](https://cdn.jsdelivr.net/gh/asiaqa/ag-rule@release/output_super_hosts.txt) |
| Full | [output_full_hosts.txt](https://raw.githubusercontent.com/asiaqa/ag-rule/main/rules/output_full_hosts.txt) | [jsdelivr](https://cdn.jsdelivr.net/gh/asiaqa/ag-rule@release/output_full_hosts.txt) |
| Medium | [output_medium_hosts.txt](https://raw.githubusercontent.com/asiaqa/ag-rule/main/rules/output_medium_hosts.txt) | [jsdelivr](https://cdn.jsdelivr.net/gh/asiaqa/ag-rule@release/output_medium_hosts.txt) |
| Min | [output_min_hosts.txt](https://raw.githubusercontent.com/asiaqa/ag-rule/main/rules/output_min_hosts.txt) | [jsdelivr](https://cdn.jsdelivr.net/gh/asiaqa/ag-rule@release/output_min_hosts.txt) |

### dnsmasq Format (OpenWrt)

| Tier | Raw URL | CDN |
|------|---------|-----|
| Super | [output_super_dnsmasq.txt](https://raw.githubusercontent.com/asiaqa/ag-rule/main/rules/output_super_dnsmasq.txt) | [jsdelivr](https://cdn.jsdelivr.net/gh/asiaqa/ag-rule@release/output_super_dnsmasq.txt) |
| Full | [output_full_dnsmasq.txt](https://raw.githubusercontent.com/asiaqa/ag-rule/main/rules/output_full_dnsmasq.txt) | [jsdelivr](https://cdn.jsdelivr.net/gh/asiaqa/ag-rule@release/output_full_dnsmasq.txt) |
| Medium | [output_medium_dnsmasq.txt](https://raw.githubusercontent.com/asiaqa/ag-rule/main/rules/output_medium_dnsmasq.txt) | [jsdelivr](https://cdn.jsdelivr.net/gh/asiaqa/ag-rule@release/output_medium_dnsmasq.txt) |
| Min | [output_min_dnsmasq.txt](https://raw.githubusercontent.com/asiaqa/ag-rule/main/rules/output_min_dnsmasq.txt) | [jsdelivr](https://cdn.jsdelivr.net/gh/asiaqa/ag-rule@release/output_min_dnsmasq.txt) |

### Unbound Format

| Tier | Raw URL | CDN |
|------|---------|-----|
| Super | [output_super_unbound.txt](https://raw.githubusercontent.com/asiaqa/ag-rule/main/rules/output_super_unbound.txt) | [jsdelivr](https://cdn.jsdelivr.net/gh/asiaqa/ag-rule@release/output_super_unbound.txt) |
| Full | [output_full_unbound.txt](https://raw.githubusercontent.com/asiaqa/ag-rule/main/rules/output_full_unbound.txt) | [jsdelivr](https://cdn.jsdelivr.net/gh/asiaqa/ag-rule@release/output_full_unbound.txt) |
| Medium | [output_medium_unbound.txt](https://raw.githubusercontent.com/asiaqa/ag-rule/main/rules/output_medium_unbound.txt) | [jsdelivr](https://cdn.jsdelivr.net/gh/asiaqa/ag-rule@release/output_medium_unbound.txt) |
| Min | [output_min_unbound.txt](https://raw.githubusercontent.com/asiaqa/ag-rule/main/rules/output_min_unbound.txt) | [jsdelivr](https://cdn.jsdelivr.net/gh/asiaqa/ag-rule@release/output_min_unbound.txt) |

### uBlock Origin Format

| Tier | Raw URL | CDN |
|------|---------|-----|
| Super | [output_super_ublock.txt](https://raw.githubusercontent.com/asiaqa/ag-rule/main/rules/output_super_ublock.txt) | [jsdelivr](https://cdn.jsdelivr.net/gh/asiaqa/ag-rule@release/output_super_ublock.txt) |
| Full | [output_full_ublock.txt](https://raw.githubusercontent.com/asiaqa/ag-rule/main/rules/output_full_ublock.txt) | [jsdelivr](https://cdn.jsdelivr.net/gh/asiaqa/ag-rule@release/output_full_ublock.txt) |
| Medium | [output_medium_ublock.txt](https://raw.githubusercontent.com/asiaqa/ag-rule/main/rules/output_medium_ublock.txt) | [jsdelivr](https://cdn.jsdelivr.net/gh/asiaqa/ag-rule@release/output_medium_ublock.txt) |
| Min | [output_min_ublock.txt](https://raw.githubusercontent.com/asiaqa/ag-rule/main/rules/output_min_ublock.txt) | [jsdelivr](https://cdn.jsdelivr.net/gh/asiaqa/ag-rule@release/output_min_ublock.txt) |

### Domain-only List

| Tier | Raw URL | CDN |
|------|---------|-----|
| Super | [output_super_domains.txt](https://raw.githubusercontent.com/asiaqa/ag-rule/main/rules/output_super_domains.txt) | [jsdelivr](https://cdn.jsdelivr.net/gh/asiaqa/ag-rule@release/output_super_domains.txt) |
| Full | [output_full_domains.txt](https://raw.githubusercontent.com/asiaqa/ag-rule/main/rules/output_full_domains.txt) | [jsdelivr](https://cdn.jsdelivr.net/gh/asiaqa/ag-rule@release/output_full_domains.txt) |
| Medium | [output_medium_domains.txt](https://raw.githubusercontent.com/asiaqa/ag-rule/main/rules/output_medium_domains.txt) | [jsdelivr](https://cdn.jsdelivr.net/gh/asiaqa/ag-rule@release/output_medium_domains.txt) |
| Min | [output_min_domains.txt](https://raw.githubusercontent.com/asiaqa/ag-rule/main/rules/output_min_domains.txt) | [jsdelivr](https://cdn.jsdelivr.net/gh/asiaqa/ag-rule@release/output_min_domains.txt) |

### Other Lists

| List | Raw URL | CDN |
|------|---------|-----|
| GitHub Hosts (China) | [hosts](https://raw.githubusercontent.com/asiaqa/ag-rule/main/rules/hosts) | [jsdelivr](https://cdn.jsdelivr.net/gh/asiaqa/ag-rule@release/hosts) |
| DoH Blocklist | [output_doh_block.txt](https://raw.githubusercontent.com/asiaqa/ag-rule/main/rules/output_doh_block.txt) | [jsdelivr](https://cdn.jsdelivr.net/gh/asiaqa/ag-rule@release/output_doh_block.txt) |
| Dead Domains | [dead_domains.txt](https://raw.githubusercontent.com/asiaqa/ag-rule/main/rules/dead_domains.txt) | - |

---

## Format Description / 格式說明

| Format | Example | Use Case |
|--------|---------|----------|
| AdGuard Home | `\|\|ads.example.com^` | AdGuard Home, AdGuard DNS |
| Hosts | `0.0.0.0 ads.example.com` | Pi-hole, `/etc/hosts` |
| dnsmasq | `address=/ads.example.com/#` | OpenWrt, dnsmasq |
| Unbound | `local-zone: "ads.example.com" NXDOMAIN` | Unbound DNS resolver |
| uBlock Origin | `\|\|ads.example.com^` | Browser extensions |
| Domain-only | `ads.example.com` | Custom scripts |

---

## DNS NXDOMAIN Filtering / DNS 過濾

Each build checks domains via DNS using [puredns](https://github.com/d3mondev/puredns). Domains returning NXDOMAIN are removed from the output to reduce file size.

每次建置時使用 puredns 檢查網域的 DNS 狀態。返回 NXDOMAIN 的網域會從輸出中移除，以減少檔案大小。

- Dead domains are listed in `rules/dead_domains.txt` for transparency
- Results are cached for 24 hours to avoid redundant lookups
- 1M+ domains can be checked in approximately 5 minutes

---

## How to Use / 使用方法

**Pi-hole (Hosts format)**
```
https://raw.githubusercontent.com/asiaqa/ag-rule/main/rules/output_full_hosts.txt
```

**dnsmasq (OpenWrt)**
```
https://raw.githubusercontent.com/asiaqa/ag-rule/main/rules/output_full_dnsmasq.txt
```

**Unbound**
```
https://raw.githubusercontent.com/asiaqa/ag-rule/main/rules/output_full_unbound.txt
```

**AdGuard Home**
```
https://raw.githubusercontent.com/asiaqa/ag-rule/main/rules/output_full.txt
```

---

## Thanks / 致謝

- [AdguardTeam/AdGuardHome](https://github.com/AdguardTeam/AdGuardHome)
- [adaway.org](https://adaway.org/hosts.txt)
- [easylist](https://easylist-downloads.adblockplus.org/easylist.txt)
- [easylistchina](https://easylist-downloads.adblockplus.org/easylistchina.txt)
- [easyprivacy](https://easylist-downloads.adblockplus.org/easyprivacy.txt)
- [hagezi/dns-blocklists](https://github.com/hagezi/dns-blocklists)
- [neodevpro/neodevhost](https://github.com/neodevpro/neodevhost)
- [521xueweihan/GitHub520](https://github.com/521xueweihan/GitHub520)
- [dibdot/DoH-IP-blocklists](https://github.com/dibdot/DoH-IP-blocklists)
- [Perflyst/PiHoleBlocklist](https://github.com/Perflyst/PiHoleBlocklist)
- [cjx82630/cjxlist](https://github.com/cjx82630/cjxlist)
- [filter.futa.gg](https://filter.futa.gg/hosts.txt)
- [XIU2/TrackersListCollection](https://github.com/XIU2/TrackersListCollection)

---

## License

MIT
