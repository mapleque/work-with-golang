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
