package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

// По условию
func workerCondition() {
	fmt.Println("workerCondition running...")
	for i := 0; i < 10; i++ {
		if i == 5 {
			fmt.Println("workerCondition finished...")
			return
		}
	}
}

// Канал
func workerChan(done <-chan bool) {
	fmt.Println("workerChan running...")
	for {
		select {
		case <-done:
			fmt.Println("workerChan finished...")
			return
		default:
		}
	}
}

// Контекст
func workerContext(ctx context.Context) {
	fmt.Println("workerContext running...")
	for {
		select {
		case <-ctx.Done():
			fmt.Println("workerContext finished...")
			return
		default:
		}
	}
}

// runtime.Goexit
func workerGoexit() {
	fmt.Println("workerGoexit running...")
	time.Sleep(1 * time.Second)
	fmt.Println("workerGoexit finished...")
	runtime.Goexit()
}

func workerPanic() {
	fmt.Println("workerPanic running...")
	time.Sleep(1 * time.Second)
	panic("workerPanic finished...")
}

func main() {
	// Завершение по условию
	go workerCondition()

	// Завершение через канал
	stop := make(chan bool)
	go workerChan(stop)
	time.Sleep(1 * time.Second)
	stop <- true

	// Завершение через контекст
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	go workerContext(ctx)
	time.Sleep(1 * time.Second)
	cancel()

	// Завершение через runtime.Goexit
	go workerGoexit()
	time.Sleep(1 * time.Second)

	// Завершение через panic
	go workerPanic()
	time.Sleep(1 * time.Second)
}
