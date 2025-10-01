package main

func main() {
	text := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

	for i := 0; i < 10000000; i++ {
		toLower(text)
	}
}

func toLower(input string) string {
	res := ""
	const distance = 'a' - 'A'
	for _, ch := range input {
		if ch >= 'A' && ch <= 'Z' {
			res += string(ch | distance)
		} else {
			res += string(ch)
		}
	}
	return res
}
