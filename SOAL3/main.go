package main

import (
	"fmt"
	"strings"
)

func main() {
	var data = "(stockbit)"
	result := findFirstStringInBracket(data)
	resultNew := findFirstStringInBracketNew(data)

	fmt.Println("sebelum refactor = ", result)
	fmt.Println("seteleah refactor = ", resultNew)
}

// in assessment
func findFirstStringInBracket(str string) string {
	if len(str) > 0 {
		indexFirstBracketFound := strings.Index(str, "(")
		if indexFirstBracketFound >= 0 {
			runes := []rune(str)
			wordsAfterFirstBracket := string(runes[indexFirstBracketFound:len(str)])
			indexClosingBracketFound := strings.Index(wordsAfterFirstBracket, ")")
			if indexClosingBracketFound >= 0 {
				runes := []rune(wordsAfterFirstBracket)
				// return string(runes[1 : indexClosingBracketFound-1])
				return string(runes[1:indexClosingBracketFound])

			} else {
				return ""
			}
		} else {
			return ""
		}
	} else {
		return ""
	}
	return ""
}

// my answer
func findFirstStringInBracketNew(s string) string {
	i := strings.IndexByte(s, '(')
	if i < 0 {
		return ""
	}
	i++
	// j := strings.IndexByte(s[i+1:], ')')
	j := strings.IndexByte(s[i:], ')')
	if j < 0 {
		return ""
	}
	return s[i : i+j]
}
