package main

import "fmt"

func main() {
	// defer is used to ensure that a function call is performed later in a program's execution, usually for purposes of cleanup.
	defer cleanUp()

	if true {
		panic("An error accured")
	}
}

func cleanUp() {
	fmt.Println("cleanUp worked")
}
