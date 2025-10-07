package main

import "fmt"

func setBit(input int64, bitIndex uint, bitValue int) int64 {
	if bitIndex < 1 || bitIndex > 64 {
		panic("Bit index must be between 1 and 64")
	}

	if bitValue == 1 {
		return input | (1<<bitIndex - 1)
	} else if bitValue == 0 {
		return input &^ (1<<bitIndex - 1)
	} else {
		panic("Bit value must be 0 or 1")
	}
}

func main() {
	var input int64 = 5

	fmt.Printf("input = %d (%08b)\n", input, input)

	result1 := setBit(input, 1, 0)
	fmt.Printf("Установка 1-го бита в 0: %d (%08b)\n", result1, result1)

	result2 := setBit(input, 2, 1)
	fmt.Printf("Установка 2-го бита в 1: %d (%08b)\n", result2, result2)
}
