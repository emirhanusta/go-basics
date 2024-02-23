package main

import "fmt"

func main() {
	/*total := sumAll(10, 20, 30, 40, 50)
	fmt.Println(total)*/
	total, difference, multiplication, division := calculate(10, 5)
	fmt.Printf("Total: %d, Difference: %d, Multiplication: %d, Division: %d \n", total, difference, multiplication, division)
}
func sumAll(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}
func calculate(x int, y int) (int, int, int, int) {
	return x + y, x - y, x * y, x / y
}
