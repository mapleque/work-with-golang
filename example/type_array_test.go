package example

import (
	"fmt"
)

func Example_array() {
	var array1 [3]int
	array2 := [2]bool{false, true}
	array3 := [...]uint{1, 2, 3}
	array4 := [5]int{1: 1, 3: 10}

	fmt.Printf("len=%d cap=%d array=%v\n", len(array1), cap(array1), array1)
	fmt.Printf("len=%d cap=%d array=%v\n", len(array2), cap(array2), array2)
	fmt.Printf("len=%d cap=%d array=%v\n", len(array3), cap(array3), array3)
	fmt.Printf("len=%d cap=%d array=%v\n", len(array4), cap(array4), array4)

	// visit element
	fmt.Println(array4[3])

	var slice1 []int
	slice2 := []bool{false, true}
	slice3 := make([]int, 3, 5)
	slice4 := array4[1:3]
	fmt.Printf("len=%d cap=%d slice=%v\n", len(slice1), cap(slice1), slice1)
	fmt.Printf("len=%d cap=%d slice=%v\n", len(slice2), cap(slice2), slice2)
	fmt.Printf("len=%d cap=%d slice=%v\n", len(slice3), cap(slice3), slice3)
	fmt.Printf("len=%d cap=%d slice=%v\n", len(slice4), cap(slice4), slice4)

	fmt.Println("----")
	copy(slice1, slice4)
	copy(slice3, slice4)
	slice5 := make([]int, 1, 1)
	copy(slice5, slice4)
	fmt.Printf("from len=%d cap=%d slice=%v\n", len(slice4), cap(slice4), slice4)
	fmt.Printf("copy to empty len=%d cap=%d slice=%v\n", len(slice1), cap(slice1), slice1)
	fmt.Printf("copy to large len=%d cap=%d slice=%v\n", len(slice3), cap(slice3), slice3)
	fmt.Printf("copy to small len=%d cap=%d slice=%v\n", len(slice5), cap(slice5), slice5)

	fmt.Println("----")
	fmt.Printf("before grow len=%d cap=%d slice=%v\n", len(slice4), cap(slice4), slice4)
	slice4 = append(slice4, 100, 101)
	fmt.Printf("not grow len=%d cap=%d slice=%v\n", len(slice4), cap(slice4), slice4)
	slice6 := slice4    // copy
	slice7 := slice4[:] // copy
	slice4 = append(slice4, 102)
	slice4[0] = 200
	fmt.Printf("auto grow len=%d cap=%d slice=%v\n", len(slice4), cap(slice4), slice4)
	fmt.Printf("copy len=%d cap=%d slice=%v\n", len(slice6), cap(slice6), slice6)
	fmt.Printf("copy len=%d cap=%d slice=%v\n", len(slice7), cap(slice7), slice7)

	// Output:
	// len=3 cap=3 array=[0 0 0]
	// len=2 cap=2 array=[false true]
	// len=3 cap=3 array=[1 2 3]
	// len=5 cap=5 array=[0 1 0 10 0]
	// 10
	// len=0 cap=0 slice=[]
	// len=2 cap=2 slice=[false true]
	// len=3 cap=5 slice=[0 0 0]
	// len=2 cap=4 slice=[1 0]
	// ----
	// from len=2 cap=4 slice=[1 0]
	// copy to empty len=0 cap=0 slice=[]
	// copy to large len=3 cap=5 slice=[1 0 0]
	// copy to small len=1 cap=1 slice=[1]
	// ----
	// before grow len=2 cap=4 slice=[1 0]
	// not grow len=4 cap=4 slice=[1 0 100 101]
	// auto grow len=5 cap=8 slice=[200 0 100 101 102]
	// copy len=4 cap=4 slice=[1 0 100 101]
	// copy len=4 cap=4 slice=[1 0 100 101]
}
