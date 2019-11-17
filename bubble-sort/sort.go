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
	for pivot := length - 1; pivot > 0; pivot-- {
		for i := 0; i < pivot; i++ {
			if array[i] > array[i+1] {
				array[i], array[i+1] = array[i+1], array[i]
			}

		}

	}

	return array
}
