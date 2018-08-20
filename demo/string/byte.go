package main

import (
	"fmt"
)

func main() {
	byteSlice := []byte{0x43, 0x61, 0x66, 0xC3, 0xA9}
	str := string(byteSlice)
	fmt.Println(str)

	byteSlice2 := []byte{67, 97, 102, 195, 169} //decimal equivalent of {'\x43', '\x61', '\x66', '\xC3', '\xA9'}
	str2 := string(byteSlice2)
	fmt.Println(str2)

	runeSlice := []rune{0x0053, 0x0065, 0x00f1, 0x006f, 0x0072}
	str3 := string(runeSlice)
	fmt.Println(str3)
}
