package main

import (
	"fmt"
	"math/big"
)

func main() {
	a, _ := new(big.Int).SetString("5000000000000000000", 10) // 5e18 > 2^20
	b, _ := new(big.Int).SetString("3000000000000000000", 10) // 3e18 > 2^20

	result := new(big.Int)

	sum := result.Add(a, b)
	fmt.Printf("%v + %v = %v\n", a, b, sum)

	sub := new(big.Int)
	fmt.Printf("%v - %v = %v\n", a, b, sub)

	mul := new(big.Int).Mul(a, b)
	fmt.Printf("%v * %v = %v\n", a, b, mul)

	div := new(big.Int).Div(a, b)
	fmt.Printf("%v / %v = %v\n", a, b, div)

	quotient := new(big.Int)
	remainder := new(big.Int)
	quotient.QuoRem(a, b, remainder)
	fmt.Printf("%v / %v = %v (остаток: %v)\n", a, b, quotient, remainder)

}
