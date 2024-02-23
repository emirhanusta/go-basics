package main

import "fmt"

func main() {
	var names [3]string
	names[0] = "John"
	names[1] = "Doe"
	names[2] = "Smith"
	var ages = [3]int{25, 30, 35}
	var scores = []int{100, 200, 300}
	scores = append(scores, 600)
	fmt.Println(names[0:2])
	fmt.Println(ages)
	fmt.Println(scores)
}
