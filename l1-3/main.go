package main

import (
	"context"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"time"
)

func worker(dataChan <-chan int, workerNum int) {
	for n := range dataChan {
		fmt.Printf("Worker %d received data: %d\n", workerNum, n)
	}
}

func main() {
	var nFlag = flag.Int("n", 10, "gorutine amount")
	flag.Parse()

	n := *nFlag

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	dataChan := make(chan int)

	for i := 0; i < n; i++ {
		go worker(dataChan, i)
	}

loop:
	for {
		select {
		case <-ctx.Done():
			break loop
		default:
			dataChan <- rand.Intn(10)
			time.Sleep(time.Millisecond * 500)
		}
	}

	close(dataChan)
	time.Sleep(time.Millisecond * 100)
	fmt.Println("Program done")
}
