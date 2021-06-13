package example

import "fmt"

func Example_func() {
	fmt.Println("step0")
	defer func() {
		fmt.Println("defer1")
	}()
	fmt.Println("step1")
	defer func() {
		fmt.Println("defer2")
	}()
	fmt.Println("step2")
	if true {
		return
	}
	defer func() {
		fmt.Println("defer3")
	}()
	fmt.Println("step3")

	// Output:
	// step0
	// step1
	// step2
	// defer2
	// defer1
}
