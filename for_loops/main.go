package main

import "fmt"

func main() {
	//for loop
	/*var names = []string{"John", "Paul", "George", "Ringo"}
	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}
	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}
		fmt.Println(i)
	}*/

	//foreach loop
	/*var numbers = []int{1, 2, 3, 4, 5}

	for _, value := range numbers {
		fmt.Println(value)
	}*/

	/*var language = "Golang"
	for _, character := range language {
		fmt.Println(string(character))
	}*/
	var person = map[string]string{
		"name":    "John",
		"address": "Los Angeles",
		"age":     "35",
	}
	for key, value := range person {
		fmt.Println(key, ":", value)
	}

}
