package main

import "fmt"

func main() {
	//明示的な宣言
	// var 変数名 型 = 値
	var i int = 100
	fmt.Println(i)

	var s string = "Hello Golang!"
	fmt.Println(s)

	// まとめて定義
	var t, f bool = true, false
	fmt.Println(t, f)

	var (
		i2 int    = 200
		s2 string = "Hello s2 string"
	)
	fmt.Println(i2, s2)

	// 型宣言のみ
	var i3 int
	var s3 string
	fmt.Println(i3, s3) // 0と''(空文字)

	//　値の代入
	i3 = 300
	s3 = "Golnag"
	fmt.Println(i3, s3)
}
