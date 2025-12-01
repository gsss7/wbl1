package main

import "fmt"

func reverseString(s string) string {
	runes := []rune(s)

	i, j := 0, len(runes)-1
	for i < j {
		runes[i], runes[j] = runes[j], runes[i]
		i++
		j--
	}
	return string(runes)
}

func main() {
	tests := []string{
		"главрыба",
		"Hello, 世界! 👋",
	}

	for _, test := range tests {
		fmt.Printf("Original: %q]\n", test)
		fmt.Printf("Original: %q\n", reverseString(test))
	}
}
