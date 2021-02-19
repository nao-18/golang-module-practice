package main

import "fmt"

func main() {
	// 無名関数
	f := func(x, y int) int {
		return x + y
	}
	i := f(1, 4)
	fmt.Println(i) // 5

	// そのまま変数で無名関数の値を受け取る
	i2 := func(x, y int) int {
		return x + y
	}(1, 3)
	fmt.Println(i2) // 4
}
