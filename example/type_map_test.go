package example

import (
	"fmt"
)

func Example_map() {
	var map1 map[int]int
	fmt.Println(map1)
	fmt.Println(map1[0])
	v1, exist := map1[0]
	fmt.Println(v1, exist)

	map2 := map[string]string{
		"val1": "hello",
		"val2": "world",
	}
	fmt.Println(map2)
	fmt.Println(map2["val1"])
	v2, exist := map2["val2"]
	fmt.Println(v2, exist)

	// Following code will usually panic cause by:
	// concurrent map writes
	//
	// for i := 0; i < 10; i++ {
	// 	go func() {
	// 		for j := 0; j < 100; j++ {
	// 			map2["val1"] = "abc"
	// 		}
	// 	}()
	// }

	// Output:
	// map[]
	// 0
	// 0 false
	// map[val1:hello val2:world]
	// hello
	// world true
}
