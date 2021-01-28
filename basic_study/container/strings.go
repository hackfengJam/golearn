package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "Yes哈哈哈！"                // UTF-8 ，可变长编码
	fmt.Println(s)                // 12
	fmt.Println(len(s))           // 12
	fmt.Printf("%X\n", []byte(s)) // 12
	for _, b := range []byte(s) {
		fmt.Printf("%X ", b)
	}
	fmt.Println()
	for i, ch := range s { // ch is a rune
		fmt.Printf("(%d %X)", i, ch)
	}
	fmt.Println()

	fmt.Println("Rune count:",
		utf8.RuneCountInString(s))

	bytes := []byte(s)
	for len(bytes) > 0 {
		ch, size := utf8.DecodeRune(bytes)
		bytes = bytes[size:]
		fmt.Printf("%c ", ch)
	}
	fmt.Println()

	for i, ch := range []rune(s) {
		fmt.Printf("(%d %c)", i, ch)
	}
	fmt.Println()
}
