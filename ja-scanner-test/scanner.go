package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
	"golang.org/x/text/unicode/norm"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter the relative path to the file: ")
	scanner.Scan()
	filepath := scanner.Text()

	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	scanner = bufio.NewScanner(file)
	lineNo := 1
	for scanner.Scan() {
		line := scanner.Text()
		if containsJapanese(line) {
			fmt.Printf("Line %d: %s\n", lineNo, line)
		}
		lineNo++
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
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
