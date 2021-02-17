package main

import "fmt"

/*
	文字列はバイト配列の集まり
*/

func main() {
	// 文字リテラルは""(ダブルクォーテーション)で囲む
	var s string = "Hello Golang"
	fmt.Println(s)        // Hello Golang
	fmt.Printf("%T\n", s) // string

	var s2 string = "300"
	fmt.Println(s2)        // 300
	fmt.Printf("%T\n", s2) // string

	// 複数行文字列は`(バッククォーテーション)を使う
	fmt.Println(`test
	test
		test`)

	fmt.Println("\"")
	fmt.Println(`"`)

	// 文字列の特定文字を表示
	fmt.Println(s[0])         // そのままだとバイト型で表示される
	fmt.Println(string(s[0])) // string()で型変換が必要
}
