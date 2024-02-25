package main

import "fmt"

func main() {

	messages := make(chan string, 2)

	go func() {
		data1 := <-messages
		fmt.Println(data1)

		data2 := <-messages
		fmt.Println(data2)
	}()

	messages <- "Hello"
	messages <- "World"
	messages <- "Another message"

	fmt.Println("Done")
}
