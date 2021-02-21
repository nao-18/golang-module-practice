package main

import "fmt"

func main() {
	sl := []int{100, 200}
	fmt.Println(sl) // [100 200]

	// append 要素の追加(最後に追加)
	sl = append(sl, 300)
	fmt.Println(sl) // [100 200 300]

	// 複数追加も可能
	sl = append(sl, 400, 500, 600)
	fmt.Println(sl) // [100 200 300 400 500 600]

	// make　スライスの生成
	sl2 := make([]int, 5)
	fmt.Println(sl2) // [0 0 0 0 0]

	// len 要素数を調べる
	fmt.Println(len(sl2)) // 5

	// cap 容量を調べる
	fmt.Println(cap(sl2)) // 5

	// make関数は第三引数で容量を設定できる
	sl3 := make([]int, 5, 10)
	fmt.Println(len(sl3)) // 5
	fmt.Println(cap(sl3)) // 10

	sl3 = append(sl3, 1, 2, 3, 4, 5, 6, 7)
	fmt.Println(len(sl3)) // 12
	fmt.Println(cap(sl3)) // 20 キャパシティが倍増されてしまうので注意
}
