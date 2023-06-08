package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <filename>")
		return
	}

	filename := os.Args[1]

	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Error opening file: %s\n", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineNumber := 1
	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Fields(line)
		coloredLine := ""
		for _, word := range words {
			if len(word) > 7 {
				coloredLine += color.RedString(word) + " "
			} else {
				coloredLine += word + " "
			}
		}
		fmt.Printf("%d: %s\n", lineNumber, coloredLine)
		lineNumber++
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %s\n", err)
		return
	}
}
