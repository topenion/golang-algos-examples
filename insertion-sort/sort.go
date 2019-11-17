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
	for i := 1; i < length; i++ {
		for j := i; j >= 1; j-- {
			if array[j] < array[j-1] {
				array[j], array[j-1] = array[j-1], array[j]
			} else {
				break //if equal or smaller no need to iterate more
			}
		}
	}
	return array
}
