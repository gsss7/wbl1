package main

import (
	"fmt"
	"sync"
)

func generateNumbers(nums []int, output chan<- int) {
	defer close(output)
	for _, n := range nums {
		output <- n
	}
}

func processNumbers(input <-chan int, output chan<- int) {
	defer close(output)
	for n := range input {
		output <- n * 2
	}
}

func main() {
	numbers := []int{2, 5, 6, 8, 4, -9, -15}

	nChan := make(chan int)
	doubleChan := make(chan int)

	var wg sync.WaitGroup
	wg.Add(1)

	// Генерация
	go generateNumbers(numbers, nChan)
	// Обработка
	go processNumbers(nChan, doubleChan)

	go func() {
		defer wg.Done()
		for n := range doubleChan {
			fmt.Println("received:", n)
		}
	}()
	wg.Wait()
	fmt.Println("Finished printing from channel")
}
