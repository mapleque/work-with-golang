package example

import (
	"fmt"
)

func Example_bool() {
	var deafultBoolParam bool
	fmt.Println(deafultBoolParam)

	var boolParam bool = true
	fmt.Println(boolParam)
	boolParam = false
	fmt.Println(boolParam)

	directBoolParam := true
	fmt.Println(directBoolParam)

	// Output:
	// false
	// true
	// false
	// true
}
