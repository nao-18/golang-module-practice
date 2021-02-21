package main

import "fmt"

/*
	オブジェクト指向のクラス継承のようなもの
*/
type Point struct {
	A int
	B string
	C float64
}

type BigPoint struct {
	// Point Point
	// フィールド名が同じ場合は省略して記入可能
	Point
	A int
}

//メソッド
func (p *Point) Set(i int) {
	p.A = i
}

func main() {
	bp := BigPoint{}
	fmt.Println(bp) // {{0  0}}

	bp.Point.A = 100
	bp.Point.B = "Apple"
	bp.Point.C = 2.09

	fmt.Println(bp) // {{100 Apple 2.09}}
	// 省略して書いた場合は直接フィールドを呼び出せる
	fmt.Println(bp.A) // 100

	//BigPointで定義したフィールドAの値を更新する場合
	bp.A = 1000
	fmt.Println(bp.A) // 1000

	// BigPointに初期値を持たせる場合
	bp2 := BigPoint{
		Point: Point{
			A: 100,
			B: "Banana",
			C: 2.9,
		},
	}
	fmt.Println(bp2) // {{100 Banana 2.9} 0}

	// メソッドでPointのフィールドの値を更新
	bp2.Point.Set(3333)
	fmt.Println(bp2) // {{3333 Banana 2.9} 0}
}
