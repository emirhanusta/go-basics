package main

func main() {
	var age = 23
	if age >= 18 {
		println("You are eligible to vote")
	} else {
		println("You are not eligible to vote")
	}

	a := 10
	b := 20
	c := 30
	if a >= b && a >= c {
		println("a is the largest number")
	} else if b >= a && b >= c {
		println("b is the largest number")
	} else {
		println("c is the largest number")
	}

}
