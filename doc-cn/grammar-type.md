基本类型
====

在Go语言中，基本类型都可以通过关键词直接声明。

布尔
----

布尔类型通过关键词`bool`声明。默认值`false`，可赋值为`true`。

```golang
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
```

整型
----

在Go语言中，`int`表示对应CPU位数长度的整形数字，
还可以通过在`int`后面加数字来指定位数，如`int8` `int16` `int32` `int64`。

```golang
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
```

无符号整型
----

与整型类似，关键词有`uint` `uint8` `uint16` `uint32` `uint64`。

```golang
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
```

数组
----

在Go语言中，使用符号`[]`放在类型关键词之前来声明数组。

特别需要注意的，在Go语言中，数组一定是一组类型一致且长度固定的数据结构。
为了使用方便，Go语言中还提供了一个与数组形式几乎一样的可变长度的数据结构，
我们称之为切片。

数组和切片在使用过程中可能会发生隐世转换，通常我们不需要特别关心。
为了避免出错，建议在可以牺牲一定性能的情况下主动使用切片。

注意下面代码中的初始化方式。

```golang
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
```

字典
----

在Go语言中，使用`map`关键词来声明字典。字典的key和value可以是任意类型。

需要特别注意的是，并发对字典进行写的时候，会产生恐慌(panic)。


```golang
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
```

字符串
----

在Go语言中，`string` `[]byte` `[]rune`三种类型都可以表示一个字符串。

其中， byte类型是uint8类型的别名，
底层通过一个uint8的整数表示Assii码表中的一个字符。


rune类型是int32类型的别名，
底层通过一个int32的整数表示一个Unicode编码的字符。

特别的，在Go语言中，所有的Unicode都默认采用UTF-8编码规则实现。

因此，string可以转化成两种形式表示：`[]byte` `[]rune`。
显然，在表示同一个string时，这两种形式的数组长度可能是不同的。


```golang
package example

import (
	"fmt"
	"unicode/utf8"
)

func Example_string() {
	var bt byte = 0
	fmt.Printf("min byte is %d\n", bt)
	fmt.Printf("max byte is %d\n", ^bt)

	var rn rune = rune(^uint32(0) >> 1)
	fmt.Printf("min rune is %d\n", ^rn)
	fmt.Printf("max rune is %d\n", rn)

	var simpleStr string = "abcd"
	fmt.Printf("simple str bytes is %v\n", []byte(simpleStr))
	fmt.Printf("simple str runes is %v\n", []rune(simpleStr))

	var complexStr string = "我爱中国"
	cpxBt := []byte(complexStr)
	cpxRn := []rune(complexStr)
	fmt.Printf("complex str bytes is %v\n", cpxBt)
	fmt.Printf("complex str runes is %v\n", cpxRn)

	if r, s := utf8.DecodeRune(cpxBt); r == cpxRn[0] && s == 3 {
		fmt.Println("here one chinese rune combine with 3 bytes in utf8 encoding")
	} else {
		fmt.Printf("%b, %d\n", r, s)
		fmt.Printf("%b\n", cpxRn[0])
		fmt.Printf("%b|%b|%b\n", cpxBt[0], cpxBt[1], cpxBt[2])
	}

	// Output:
	// min byte is 0
	// max byte is 255
	// min rune is -2147483648
	// max rune is 2147483647
	// simple str bytes is [97 98 99 100]
	// simple str runes is [97 98 99 100]
	// complex str bytes is [230 136 145 231 136 177 228 184 173 229 155 189]
	// complex str runes is [25105 29233 20013 22269]
	// here one chinese rune combine with 3 bytes in utf8 encoding
}
```

浮点数
----

浮点数通过关键词`float32` `float64`声明。

```golang
package example

import (
	"fmt"
	"math"
)

func Example_float() {
	var maxFloat32 float32 = math.MaxFloat32
	var minFloat32 float32 = math.SmallestNonzeroFloat32
	var maxFloat64 float64 = math.MaxFloat64
	var minFloat64 float64 = math.SmallestNonzeroFloat64
	fmt.Printf("max float32: %v\n", maxFloat32)
	fmt.Printf("min float32: %v\n", minFloat32)
	fmt.Printf("max float64: %v\n", maxFloat64)
	fmt.Printf("min float64: %v\n", minFloat64)
	pi := math.Pi
	fmt.Printf("pi: %v\n", pi)
	var one = 1.0
	fmt.Printf("one: %v\n", one)
	fmt.Printf("one+one: %v\n", one+one)
	fmt.Printf("one*one: %v\n", one*one)
	fmt.Printf("one==one: %v\n", one == one)

	// Output:
	// max float32: 3.4028235e+38
	// min float32: 1e-45
	// max float64: 1.7976931348623157e+308
	// min float64: 5e-324
	// pi: 3.141592653589793
	// one: 1
	// one+one: 2
	// one*one: 1
	// one==one: true
}
```

复数
----

在Go语言中，为复数特别定义了基本类型，
可以通过关键词`complex64` `complex128`声明。

其中，complex64的实部和虚部都是float32类型的数字，
而complex128的实部和虚部都是float64类型的数字。

```golang
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
```

