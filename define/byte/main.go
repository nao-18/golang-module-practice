package main

import "fmt"

/*
	byteはuint8(unsigned int8)
*/

func main() {
	// []はスライス
	byteA := []byte{72, 73}
	fmt.Println(byteA)        // [72 73]
	fmt.Printf("%T\n", byteA) // []uint8

	// 文字列への変換
	fmt.Println(string(byteA)) // HI

	// byteへの変換
	c := []byte("HI")
	fmt.Println(c)        // [72 73]
	fmt.Printf("%T\n", c) // []uint8
}
