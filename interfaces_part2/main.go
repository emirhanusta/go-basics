package main

import "fmt"

type IEncoder interface {
	Encode(value string)
}

type XEncoder struct {
}

type YEncoder struct {
}

func (xEncoder *XEncoder) Encode(value string) {
	fmt.Println("Encoded by XEncoder")
}
func (yEncoder *YEncoder) Encode(value string) {
	fmt.Println("Encoded by YEncoder")
}
func main() {
	var encoder IEncoder = &YEncoder{}

	encoder.Encode("23453253")
}
