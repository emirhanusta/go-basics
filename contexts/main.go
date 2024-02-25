package main

import (
	"context"
	"fmt"
)

type Product struct {
	Id   int
	Name string
}

var productChan = make(chan Product)

func main() {
	/*ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	go getProduct(1)
	select {
	case product := <-productChan:
		fmt.Println(product)
	case <-ctx.Done():
		fmt.Println("Timeout")
	}
	fmt.Println("Done")*/

	ctx := context.Background()
	ctx = context.WithValue(ctx, "key", "value")

	f1(ctx)
}
func f1(ctx context.Context) {
	fmt.Println("f1", ctx.Value("key"))
	f2(ctx)
}
func f2(ctx context.Context) {
	fmt.Println("f2", ctx.Value("key"))
	f3(ctx)
}
func f3(ctx context.Context) {
	fmt.Println("f3", ctx.Value("key"))
}

/*func getProduct(id int) {
	time.Sleep(2 * time.Second)

	productChan <- Product{Id: id, Name: "Product Name"}
}*/
