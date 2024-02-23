package main

import "fmt"

func main() {
	var d int = 4
	var b *int = &d
	var c int = *b
	fmt.Println("a:", d)
	fmt.Println("b:", b)
	fmt.Println("c:", c)

	var a int = 4

	fmt.Println(a)
	add10(&a)
	fmt.Println(a)
	var numbers = []int{2, 3, 4, 5}
	fmt.Println("numbers[0]:", numbers[0])
	changeValue(numbers)
	fmt.Println("numbers[0]:", numbers[0])
}
func changeValue(x []int) {
	x[0] = 100
}
func add10(x *int) {
	*x += 10
}
