package main

import "fmt"

func main() {
	a := 10
	b := 30

	fmt.Printf("a = %d, b = %d\n", a, b)

	a = a - b
	b = b + a
	a = b - a
	fmt.Printf("a = %d, b = %d\n", a, b)

	a = a ^ b
	b = a ^ b
	a = a ^ b
	fmt.Printf("a = %d, b = %d\n", a, b)
}
