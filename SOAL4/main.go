package main

import (
	"fmt"
	"sort"
	"strings"
)

var (
	dictionary = []string{"kita", "atik", "tika", "aku", "kia", "makan", "kua"}
)

func Anagram() {
	list := make(map[string][]string)
	var groupAnagram [][]string

	for _, word := range dictionary {
		key := sortStr(word)
		list[key] = append(list[key], word)
	}

	for _, words := range list {
		groupAnagram = append(groupAnagram, words)
	}

	fmt.Print(groupAnagram)
}

func sortStr(k string) string {
	s := strings.Split(k, "")
	sort.Strings(s)

	return strings.Join(s, "")
}

func main() {
	Anagram()
}
