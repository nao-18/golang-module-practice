package main

import "fmt"

func main() {
	/*
		sl := []int{100, 200}
		// 参照型はこう書くと参照渡しになる
		sl2 := sl
		sl2[0] = 1000
		// 参照渡しだと元のスライスも値が変わる
		fmt.Println(sl)

		// 基本型は値渡しになるため元のスライスは変わらない
		var i int = 10
		i2 := i
		i2 = 100
		fmt.Println(i, i2)
	*/

	// 参照型の変数へのコピーはcopy関数を使う
	sl := []int{1, 2, 3, 4, 5}
	sl2 := make([]int, 5, 10)
	fmt.Println(sl2) // [0 0 0 0 0 ]

	// copy関数 第一引数: コピー先, 第二引数: コピー元 戻り値: コピーに成功した要素数
	n := copy(sl2, sl)
	fmt.Println(n, sl) // 5 [1 2 3 4 5]
}
