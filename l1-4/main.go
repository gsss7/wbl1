package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func workWithContext(wg *sync.WaitGroup, ctx context.Context, workerNum int) {
	if wg != nil {
		defer wg.Done()
	}
	fmt.Printf("ctx worker %d starting...\n", workerNum)
	fmt.Printf("ctx worker %d done.\n", workerNum)

	for {
		select {
		case <-ctx.Done():
			fmt.Printf("ctx worker %d recieved cancellation signal. Clearning up ..\n", workerNum)
			time.Sleep(100 * time.Millisecond)
			fmt.Printf("ctx worker %d done.\n", workerNum)
			return
		default:
			fmt.Printf("worker %d is working...\n", workerNum)
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var wg sync.WaitGroup
	numWorkers := 5

	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go workWithContext(&wg, ctx, i)
	}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	fmt.Println("Program started. Press Ctrl+C to stop.")

	select {
	case sig := <-sigChan:
		fmt.Printf("\nReceived signal: %s. Shutting down...\n", sig)
		cancel()
	case <-ctx.Done():
	}
	wg.Wait()
	fmt.Println("All workers completed.")
}
