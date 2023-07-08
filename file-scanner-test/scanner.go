package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"unicode"
	"golang.org/x/text/unicode/norm"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter the relative path to the folder: ")
	scanner.Scan()
	root := scanner.Text()

	outputFile, err := os.Create("output.txt")
	if err != nil {
		fmt.Println("Failed to create the output file: ", err)
		return
	}
	defer outputFile.Close()

	writer := bufio.NewWriter(outputFile)

	err = filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return fmt.Errorf("Prevent panic by handling failure accessing a path %q: %v", path, err)
		}

		if !info.IsDir() {
			file, err := os.Open(path)
			if err != nil {
				return fmt.Errorf("Failed to open the file %s: %v", path, err)
			}
			defer file.Close()

			scanner = bufio.NewScanner(file)
			lineNo := 1
			for scanner.Scan() {
				line := scanner.Text()
				if containsJapanese(line) {
					fmt.Fprintf(writer, "File: %s, Line %d: %s\n", path, lineNo, line)
				}
				lineNo++
			}

			if err := scanner.Err(); err != nil {
				return fmt.Errorf("Failed to scan the file %s: %v", path, err)
			}
		}
		return nil
	})

	if err != nil {
		fmt.Fprintln(writer, err)
	}

	writer.Flush()
}

func containsJapanese(s string) bool {
	for _, r := range norm.NFKC.String(s) {
		if unicode.Is(unicode.Han, r) || // Kanji
			unicode.Is(unicode.Hiragana, r) || // Hiragana
			unicode.Is(unicode.Katakana, r) { // Katakana
			return true
		}
	}
	return false
}
