package main

import (
	"strings"
)

func main() {
	text := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

	for i := 0; i < 10000000; i++ {
		toLower(text)
	}
}

func toLower(input string) string {
	var b strings.Builder
	const distance = 'a' - 'A'
	for _, ch := range input {
		if ch >= 'A' && ch <= 'Z' {
			b.WriteRune(ch | distance)
		} else {
			b.WriteRune(ch)
		}
	}
	return b.String()
}
