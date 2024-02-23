package main

import "fmt"

type IShippable interface {
	ShippingCost() int
}

type Book struct {
	desi int
}
type Electronic struct {
	desi int
}
type Shoe struct {
	desi int
}

func (book *Book) ShippingCost() int {
	return 5 + book.desi*2
}
func (electronic *Electronic) ShippingCost() int {
	return 10 + electronic.desi*5
}
func (shoe *Shoe) ShippingCost() int {
	return 5 + shoe.desi*2
}
func main() {
	var products []IShippable = []IShippable{
		&Book{1},
		&Electronic{1},
		&Shoe{1},
	}
	fmt.Println(calculateShippingCost(products))
}
func calculateShippingCost(products []IShippable) int {
	var total int
	for _, product := range products {
		total += product.ShippingCost()
	}
	return total
}
