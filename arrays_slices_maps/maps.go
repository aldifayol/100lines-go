package array_slice_maps

import "fmt"

func MyMaps() {
	var myMap = map[string]string {
		"name" : "aldi",
		"age": "37",
	} 

	fmt.Printf("\n")
	for k, v := range myMap {
		fmt.Println(k, v)
	}
}