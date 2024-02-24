package main

import (
	"fmt"
)

func main() {

	/*fileData, err := os.ReadFile("sample.txt")

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(fileData)
	}*/

	/*result, err := divide(6, 0)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}*/

	productService := ProductService{}
	product := Product{id: 1, name: "Laptop", price: 10}
	err := productService.add(product)
	if err != nil {
		fmt.Println(err)
	}
}

type Product struct {
	id    int
	name  string
	price int
}

type ProductService struct{}

func (productService *ProductService) add(product Product) error {
	if len(product.name) == 0 {
		return ValidationError{"Product name is required", "1001"}
	}
	if product.price <= 0 {
		return ValidationError{"Product price must be greater than zero", "1002"}
	}
	fmt.Println("Product added successfully")
	return nil
}

type ValidationError struct {
	code    string
	message string
}

func (validationError ValidationError) Error() string {
	return fmt.Sprintf("Error: %s, Code: %s", validationError.message, validationError.code)
}

/*func divide(x int, y int) (int, error) {
	if y == 0 {
		return 0, errors.New("the denominator cannot be zero")
	} else {
		return x / y, nil
	}
}*/
