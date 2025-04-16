package array_slice_maps

import "fmt"

func MyArray() {
	var myArray1 = [4]int{23, 5, 6, 8}

	fmt.Println(myArray1)

	myArray2 := [3]string{
		"I want to be yours",
		"If you like",
		"Please consider to accept my feelings",
	}

	for index, option := range myArray2 {
		fmt.Println(index, option)
	}

	fmt.Println("\narray quotes: \n", myArray2[1])
}