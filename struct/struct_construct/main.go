package main

import "fmt"

type Point struct {
	A int
	B string
	C float64
}

// コンストラクタはNewから始まるメソッドで行う
// 返り値の型はポインタ型を指定する
func NewPoint(a int, b string, c float64) *Point {
	// 構造体Pointのポインタを返す
	return &Point{a, b, c}
}

func main() {
	// コンストラクタを使用しない場合
	p1 := Point{1, "A", 1.1}
	fmt.Println(p1) // {1 A 1.1}

	// コンストラクタを使用する場合
	p2 := NewPoint(1, "B", 2.35)
	fmt.Println(p2) // &{1 B 2.35}
}
