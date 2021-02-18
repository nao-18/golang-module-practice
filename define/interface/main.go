package main

import "fmt"

/*
	interfaceは全ての型と互換性がある
	初期値はnil
*/

func main() {
	// 明示的な宣言
	var x interface{}
	fmt.Println(x)        // nil
	fmt.Printf("%T\n", x) // nil

	// 値の代入はどんな型でも可能
	x = 1
	fmt.Println(x)        // 1
	fmt.Printf("%T\n", x) // int

	x = 3.4
	fmt.Println(x)        // 3.4
	fmt.Printf("%T\n", x) // float64

	x = "A"
	fmt.Println(x)        // A
	fmt.Printf("%T\n", x) // string

	x = [3]int{1, 2, 3}
	fmt.Println(x)        // [1 2 3]
	fmt.Printf("%T\n", x) // [3]int

	// 演算等の各型特有の操作はできない
	// fmt.Println(len(x))　//エラー
	// x = 2
	// fmt.Println(x + 3) //エラー
}
