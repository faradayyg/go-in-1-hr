package main

import (
	"errors"
	"fmt"
)

func main() {
	yeah, err := foo(2)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(yeah)
	}

	// Arrays
	var arr [3]int32 = [...]int32{0, 1, 2}

	// Stored in contiguous memory locations
	fmt.Println(&arr[0])
	fmt.Println(&arr[1])
	fmt.Println(&arr[2])
	fmt.Println(arr)

	// Slices: A slice does not define the length and can be increased in size
	var intSlice []int32 = []int32{1, 2, 4, 5, 6}
	fmt.Printf("The length of the slice is %v and the capacity is %v \n", len(intSlice), cap(intSlice))

	// Appending a slice to an existing slice
	intSlice = append(intSlice, []int32{3, 7, 8, 9}...)
	fmt.Printf("New length %v, new cap %v", len(intSlice), cap(intSlice))

	// Appending a single element to a slice
	intSlice = append(intSlice, 10)
	fmt.Printf("\nNew length %v, new cap(should be unchanged) %v", len(intSlice), cap(intSlice))

	// Slices using the Make function
	// make -> type, length and capacity of slice
	var intSlice2 []int32 = make([]int32, 4, 4)

	fmt.Printf("\nSize of slice %v, capacity %v", len(intSlice2), cap(intSlice2))

	//Usual operations should work as well
	intSlice2 = append(intSlice2, 5)
	intSlice2 = append(intSlice2, make([]int32, 2)...)

	fmt.Printf("\nNew size of slice %v, capacity %v", len(intSlice2), cap(intSlice2))
}

func foo(num int) (string, error) {
	var err error
	if num%2 != 0 {
		err = errors.New("Something happened, cant proceed")
	}
	return "yeah man", err
}
