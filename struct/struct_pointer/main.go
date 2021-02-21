package main

import "fmt"

type Point struct {
	A int
	B string
	C float64
}

func Update(p Point) {
	p.A = 100
	p.B = "Update"
	p.C = 4.25
}

// 引数にポインタを使う場合
func Update2(p *Point) {
	p.A = 100
	p.B = "Update"
	p.C = 4.25
}

func main() {
	p := Point{}
	Update(p)
	// 構造体は値渡しのため、元の変数pは更新されない
	fmt.Println(p) // {0  0} (0 空文字 0)

	//　アドレスを生成(推奨)
	p2 := &Point{}
	Update2(p2)
	fmt.Println(p2) // &{100 Update 4.25}

	// アドレスを直接引数に設定する場合
	Update2(&p)
	fmt.Println(p) // {100 Update 4.25}

	// アドレスを生成するnew関数を使う場合
	p3 := new(Point)
	Update2(p3)
	fmt.Println(p3) // &{100 Update 4.25}
}
