package main

import "fmt"

func determineType(v interface{}) string {
	switch v.(type) {
	case int:
		return "int"
	case string:
		return "string"
	case bool:
		return "bool"
	case chan int, chan string, chan bool, chan interface{}:
		return "chan"
	default:
		return "unknown"
	}
}

func main() {
	variables := []interface{}{
		3231,
		"str",
		true,
		make(chan int),
		make(chan string),
		make(chan bool),
		make(chan interface{}),
		3.14, // нераспознаваемый тип
	}

	fmt.Println("Базовое определение типов:")
	for _, v := range variables {
		fmt.Printf("Value: %v, Type: %s\n", v, determineType(v))
	}

}
