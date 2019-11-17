package randomgen

import (
	crand "crypto/rand"
	"encoding/binary"
	"math/rand"
)

//IntegerArray generates a random Array
func IntegerArray() []int {
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
