package control_flow

import "fmt"

func Pointerrr() {
	number := 42
	address := &number
	value := *address

	fmt.Printf("value: %v\n", value)
	fmt.Printf("address: %v\n\n", address)
}