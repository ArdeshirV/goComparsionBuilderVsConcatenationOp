package main

import (
	"strings"
	"testing"
)

func Concat(a, b string) string {
	return a + b
}

func BenchmarkConcat(b *testing.B) {
	b.Run("Concat", func(_ *testing.B) {
		_ = Concat("Hello", " World")
	})
}

func BenchmarkBuilder(b *testing.B) {
	b.Run("string.Builder", func(_ *testing.B) {
		var sb strings.Builder
		sb.WriteString("Hello,")
		sb.WriteString(" World")
		_ = sb.String()
	})
}
