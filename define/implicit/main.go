package main

import "fmt"

/*
	基本的には明示的な定義を使う
*/

// 関数外での暗黙的な定義はできない
// i2 := 3000 エラー

func main() {
	// 暗黙的宣言
	i := 400
	fmt.Println(i) // 400

	i = 40000
	fmt.Println(i) // 40000

	// 再定義不可
	// i := 5000
	// fmt.Println() エラー

	// 型推論で型は固定される
	// i = "Hello"
	// fmt.Println(i)

	// 関数外での暗黙的な定義はできない(呼び出し元)
	// fmt.Println(i2) エラー
}
