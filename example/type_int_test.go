package example

import (
	"fmt"
)

func Example_int() {
	int8max := int(^uint8(0) >> 1)
	int8min := ^int8max
	fmt.Printf("int8max = %d\n", int8max)
	fmt.Printf("int8min = %d\n", int8min)

	int16max := int(^uint16(0) >> 1)
	int16min := ^int16max
	fmt.Printf("int16max = %d\n", int16max)
	fmt.Printf("int16min = %d\n", int16min)

	int32max := int(^uint32(0) >> 1)
	int32min := ^int32max
	fmt.Printf("int32max = %d\n", int32max)
	fmt.Printf("int32min = %d\n", int32min)

	int64max := int(^uint64(0) >> 1)
	int64min := ^int64max
	fmt.Printf("int64max = %d\n", int64max)
	fmt.Printf("int64min = %d\n", int64min)

	// Output:
	// int8max = 127
	// int8min = -128
	// int16max = 32767
	// int16min = -32768
	// int32max = 2147483647
	// int32min = -2147483648
	// int64max = 9223372036854775807
	// int64min = -9223372036854775808
}
