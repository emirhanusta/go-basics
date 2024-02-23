package main

import "fmt"

func main() {

	var productName string
	var quantity int
	var discount float32
	var isInStock bool = true

	productName = "logitech"
	quantity = 1500
	discount = 0.25

	/*
		fmt.Println(productName, reflect.TypeOf(productName))
		fmt.Println(quantity, reflect.TypeOf(quantity))
		fmt.Println(discount, reflect.TypeOf(discount))
		fmt.Println(isInStock, reflect.TypeOf(isInStock))
	*/
	fmt.Printf("Product name: %s, Quantity: %d, Discount: %f, Is in stock: %t\n", productName, quantity, discount, isInStock)

	/*productName := "samsung"
	fmt.Println(productName, reflect.TypeOf(productName))*/

}
