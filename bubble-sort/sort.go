package main

import (
	crand "crypto/rand"
	"encoding/binary"
	"fmt"
	"math/rand"
)

func main() {
	array := createRandomArray()
	fmt.Printf("Unsorted Array is : %v\n", array)
	fmt.Printf("Sorted Array is : %v\n", sort(array))
}

func createRandomArray() []int {
	numOfItems := rand.Intn(8) + 2
	array := make([]int, 0)
	for i := 0; i < numOfItems; i++ {
		array = append(array, rand.Intn(20))
	}
	return array
}

func init() {
	var b [8]byte
	_, err := crand.Read(b[:])
	if err != nil {
		panic("cannot seed math/rand package with cryptographically secure random number generator")
	}
	rand.Seed(int64(binary.LittleEndian.Uint64(b[:])))
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
