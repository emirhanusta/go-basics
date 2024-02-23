package main

import "fmt"

func main() {
	/*var names map[string]int

	names = make(map[string]int, 0)

	names["John"] = 41
	names["Paul"] = 39
	names["George"] = 58

	fmt.Println(names["John"])*/
	names := map[string]int{
		"John":   41,
		"Paul":   39,
		"George": 58,
	}
	delete(names, "John")
	fmt.Println(names)
}
