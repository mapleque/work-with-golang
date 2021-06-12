package example

import (
	"fmt"
)

func Example_complext() {
	var cpx1 complex64 = complex(1, 2)
	fmt.Println(cpx1)
	fmt.Println(real(cpx1), imag(cpx1))
	fmt.Println(cpx1 + cpx1)
	fmt.Println(cpx1 * cpx1)

	// Output:
	// (1+2i)
	// 1 2
	// (2+4i)
	// (-3+4i)
}
