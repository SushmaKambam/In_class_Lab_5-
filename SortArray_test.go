package main

import (
	"fmt"
	"testing"
)

func TestSortArray(t *testing.T) {

	//Declaring an array
	arr1 := []int{5, 2, 9, 1, 5, 6}

	//Priniting original array
	fmt.Println("Original arr1:", arr1)
	//Sorting Array
	SortArray(arr1)
	//Printing the sorted array
	fmt.Println("Sorted arr1:", arr1)

}
