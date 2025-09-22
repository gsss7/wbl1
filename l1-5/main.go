package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var nFlag = flag.Int("n", 2, "ex sec")
	flag.Parse()

	N := *nFlag
	inputChan := make(chan int)
	timeout := time.After(time.Duration(N) * time.Second)

	go func() {
		for n := range inputChan {
			fmt.Printf("Received: %d\n", n)
		}
		fmt.Println("Reader: Chanel closed.")
	}()

	fmt.Println("Begin:")
	ticker := time.NewTicker(200 * time.Millisecond)
	defer ticker.Stop()

	for {
		select {
		case <-timeout:
			fmt.Println("Time is up")
			close(inputChan)
			fmt.Println("End")
			return
		case t := <-ticker.C:
			fmt.Printf("waiting, %.2fs\t", time.Since(t.Add(-200*time.Millisecond)).Seconds())
			inputChan <- rand.Intn(10)
		}
	}
}
