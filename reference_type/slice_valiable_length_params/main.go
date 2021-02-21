package main

import "fmt"

// "..."で可変長スライスを表す
func Sum(s ...int) int {
	n := 0
	for _, v := range s {
		n += v
	}
	return n
}

func main() {
	// 呼び出し時の引数はスライスになる
	fmt.Println(Sum(1, 2, 3)) // 6

	fmt.Println(Sum()) // 0
	sl := []int{2, 3, 5}
	// スライスを渡す場合は"..."をつける
	fmt.Println(Sum(sl...)) // 10
}
