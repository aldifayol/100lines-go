package main

import (
	"fmt"

	array_slice_maps "stackup.dev/intro-to-golang/arrays_slices_maps"
	"stackup.dev/intro-to-golang/control_flow"
	structsinterfaces "stackup.dev/intro-to-golang/structs_interfaces"
)

func main() {
	x := []float64{2.15, 3.14, 42.0, 29.11}
	// control_flow.HitungRata2(x)
	fmt.Println(control_flow.HitungRata2Switch(x)) // 19.1
	fmt.Println(control_flow.HitungRata2IfElse(x)) // 19.1
	fmt.Println(control_flow.HitungPakeWhile(5))   // 8

	control_flow.Pointerrr()
	structsinterfaces.Structs()
	structsinterfaces.AdvanceStruct()
	control_flow.PrintEmployee()
	array_slice_maps.MyArray()
	array_slice_maps.MySlices(1)
	array_slice_maps.MyMaps()
}
