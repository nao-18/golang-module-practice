package main

import "fmt"

// 頭文字を大文字にするとpublicになる
const Pi = 3.14

// 複数定義
const (
	URL      = "http://google.com"
	SiteName = "Google"
)

//　値がないものは前の値が入る
const (
	A = 1
	B
	C = "A"
	D
	E
)

// 整数の連番を定義数
const (
	c0 = iota
	c1
	c2
)

// 定数だと型の最大値を超えても定義できる
// var Big int = 9223372036854775807 + 1
const Big = 9223372036854775807 + 1

func main() {
	fmt.Println(Pi)

	fmt.Println(URL)
	fmt.Println(SiteName)

	fmt.Println(A, B, C, D, E)

	// fmt.Println(Big-1)
	fmt.Println(Big - 1)

	fmt.Println(c0, c1, c2)
}
