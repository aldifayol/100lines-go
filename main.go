package main

import (
	"fmt"

	"stackup.dev/intro-to-golang/control_flow"
)


func main() {
	x := []float64{2.15, 3.14, 42.0, 29.11}
	// control_flow.HitungRata2(x)
	fmt.Println(control_flow.HitungRata2Switch(x)) // 19.1
	fmt.Println(control_flow.HitungRata2IfElse(x)) // 19.1
	fmt.Println(control_flow.HitungPakeWhile(5)) // 8

	control_flow.Pointerrr()
	control_flow.Structs()
	control_flow.AdvanceStruct()
	control_flow.PrintEmployee()
}
