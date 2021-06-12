package example

import (
	"fmt"
)

func Example_uint() {
	uint8min := uint8(0)
	uint8max := ^uint8min
	fmt.Printf("uint8min = %d\n", uint8min)
	fmt.Printf("uint8max = %d\n", uint8max)

	uint16min := uint16(0)
	uint16max := ^uint16min
	fmt.Printf("uint16min = %d\n", uint16min)
	fmt.Printf("uint16max = %d\n", uint16max)

	uint32min := uint32(0)
	uint32max := ^uint32min
	fmt.Printf("uint32min = %d\n", uint32min)
	fmt.Printf("uint32max = %d\n", uint32max)

	uint64min := uint64(0)
	uint64max := ^uint64min
	fmt.Printf("uint64min = %d\n", uint64min)
	fmt.Printf("uint64max = %d\n", uint64max)

	// Output:
	// uint8min = 0
	// uint8max = 255
	// uint16min = 0
	// uint16max = 65535
	// uint32min = 0
	// uint32max = 4294967295
	// uint64min = 0
	// uint64max = 18446744073709551615
}
