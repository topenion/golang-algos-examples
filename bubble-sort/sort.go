package main

import (
	"fmt"

	randomgen "github.com/topenion/golang-algos-examples/random-generator"
)

func main() {
	array := randomgen.IntegerArray()
	fmt.Printf("Unsorted Array is : %v\n", array)
	fmt.Printf("Sorted Array is : %v\n", sort(array))
}

func sort(array []int) []int {
	length := len(array)
	for sorted := 0; sorted < length-1; sorted++ {
		for i := length - 1; i > sorted; i-- {
			if array[i-1] > array[i] {
				array[i-1], array[i] = array[i], array[i-1]
			}

		}
		fmt.Printf("Array after iteration %v is %v\n", sorted+1, array)

	}

	return array
}
