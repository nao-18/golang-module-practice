package main

import "fmt"

func main() {
	// 環境によってintは変わる
	var i int = 100

	/*
		intの種類
		int8
		int16
		int32
		int64
	*/
	fmt.Println(i)

	// 加算
	fmt.Println(i + 100)

	// 明示した型で別々に扱われる
	// var i2 int64 = 100
	// fmt.Println(i + i2)

	// 型調べ方
	fmt.Printf("%T\n", i)

	// 型変換
	fmt.Printf("%T\n", int32(i))
}
