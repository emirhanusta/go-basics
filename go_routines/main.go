package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}

	startTime := time.Now()

	wg.Add(2)

	go func() {
		defer wg.Done()
		fmt.Println("Hello")
	}()

	go func() {
		defer wg.Done()
		fmt.Println("World")
	}()

	wg.Wait()

	fmt.Println("Done")

	fmt.Println(time.Since(startTime))
}
