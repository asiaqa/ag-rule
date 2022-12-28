package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"strconv"
)

func readlink(master string, loc string) {
	inputFile, err := os.Open(master)
	if err != nil {
		fmt.Println(err)
	}
	defer inputFile.Close()
	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		// Get the current line
		download(scanner.Text(), loc)

	}
	return
}

func download(link string, loc string) {
	// Make an HTTP GET request to the URL of the file you want to download
	response, err := http.Get(link)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("download:" + link + " completed!")
	}
	defer response.Body.Close()

	// Create a local file to save the downloaded file
	file, err := os.OpenFile(loc, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	// Use io.Copy to copy the response body to the local file
	_, err = io.Copy(file, response.Body)
	if err != nil {
		fmt.Println(err)
	}
	return
}
func processing(loc string, output string, custom string) {
	i := 0 // Number of duplicated items
	j := 0 // Number of valid items
	outputFile, err := os.Create(output)
	if err != nil {
		fmt.Println(err)
	}
	linesWritten := make(map[string]bool)
	inputFile, err := os.Open(loc)
	if err != nil {
		fmt.Println(err)
	}
	defer inputFile.Close()
	scanner := bufio.NewScanner(inputFile)
	println("Creating " + output + " ......")
	for scanner.Scan() {
		// Get the current line
		line := scanner.Text()
		// Check if the line is started from "[blank] or [space] or # or !"
		match, _ := regexp.MatchString("^[ #!]|^$|^\\[", line)
		if match {
			linesWritten[line] = true
		}
		if !linesWritten[line] {
			// Write the line to the output file
			outputFile.WriteString(line + "\n")
			j++
			// Mark the line as written
			linesWritten[line] = true
		} else {
			i++
		}
	}
	customFile, err := os.Open(custom)
	if err != nil {
		fmt.Println(err)
	}
	defer customFile.Close()
	scanner1 := bufio.NewScanner(customFile)
	for scanner1.Scan() {
		outputFile.WriteString(scanner1.Text() + "\n")
		j++
	}
	defer outputFile.Close()
	println(strconv.Itoa(i) + " of items are deleted")
	println(strconv.Itoa(j) + " of items are included")
	println(output + " is created")
	return

}
func remove(loc string) {
	err := os.Remove(loc)
	if err != nil {
		fmt.Println(err)
	}
	return
}

func main() {
	readlink(os.Args[1], "temp.txt")
	processing("temp.txt", os.Args[2], os.Args[3])
	remove("temp.txt")
       // readlink("simple.txt", "temp.txt")
 //	processing("temp.txt", "out_simple.txt", "custom.txt")
//	remove("temp.txt")
}
