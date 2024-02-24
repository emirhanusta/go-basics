package main

func main() {

	ch := make(chan int)

	go func() {
		ch <- 55
	}()

	message, isOpen := <-ch

	println(message, isOpen)
}
