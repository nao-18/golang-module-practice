package main

import "fmt"

// 構造体の定義
type Point struct {
	A int
	B string
	C float64
}

func main() {
	// 明示的な宣言
	var p Point
	fmt.Println(p) // {0  0}  (0 空文字 0)

	// 暗黙的な宣言
	p2 := Point{}
	fmt.Println(p2) // {0  0}  (0 空文字 0)

	// 初期値設定
	p3 := Point{A: 1, B: "Golang", C: 1.2}
	fmt.Println(p3) // {1 Golang 1.2}

	// 値の出力
	fmt.Println(p3.A, p3.B) // 1 Golang

	// 値の更新
	p2.A = 10
	fmt.Println(p2.A) // 10

	// 複合リテラル
	// 構造体の値順で引数に設定することでkeyを省ける
	p4 := Point{1, "python", 3.14}
	fmt.Println(p4) // {1 python 3.14}

	// 構造体Point.Aのみ設定すると、B,Cは初期値になる
	p5 := Point{A: 1}
	fmt.Println(p5) // {1  0} (1 空文字 0)
}
