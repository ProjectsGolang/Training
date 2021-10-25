package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "Jorge Luis Alonso Hernandez"
	fmt.Println(WordCount(s))
}

func WordCount(text string) map[string]int {
	m := make(map[string]int)

	text_without_space := strings.ReplaceAll(text, " ", "")
	array_letters := strings.Split(text_without_space, "")

	for _, letter := range array_letters {
		if _, ok := m[letter]; ok {
			m[letter] += 1
		} else {
			m[letter] += 1
		}
	}

	return m
}
