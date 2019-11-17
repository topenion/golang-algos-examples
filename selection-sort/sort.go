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

	for pivot := 0; pivot < length-1; pivot++ {
		min := array[pivot]
		index := pivot
		for i := pivot + 1; i < length; i++ {
			if array[i] < min {
				min = array[i]
				index = i
			}
		}

		array[pivot], array[index] = array[index], array[pivot]

	}

	return array
}
