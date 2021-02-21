package main

import "fmt"

type Point struct {
	A int
	B string
	C float64
}

// 関数で値を更新した場合
func Set(p *Point, i int) {
	p.A = i
}

/*
	メソッドのレシーバーは必ず構造体のポインタ型を使用する
*/
// メソッドを使う(推奨)
func (p *Point) Set(i int) {
	p.A = i
}

// 値渡しのメソッドでは更新できない
func (p Point) Set2(i int) {
	p.A = i
}

func main() {
	p1 := &Point{A: 1}
	Set(p1, 2)
	fmt.Println(p1)  // &{2  0} アドレスを示す
	fmt.Println(*p1) // {2  0} 実体を示す

	p1.Set(5)
	fmt.Println(p1)  // &{5  0} アドレス
	fmt.Println(*p1) // {5  0} 実体

	p2 := Point{}
	p2.Set2(1000)
	fmt.Println(p2) // {0  0}

	p2.Set(99999)
	fmt.Println(p2) // {99999  0}
}
