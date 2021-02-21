package main

import "fmt"

func main() {
	// 変数の局所性を高めたい場合
	switch n := 2; n {
	case 1, 2:
		fmt.Println("1 or 2") // 1 or 2
	case 3, 4:
		fmt.Println("3 or 4")
	default:
		fmt.Println("I don't Know")
	}

	// // 列挙型と比較演算子はこのようにcaseへ追加できない
	// switch n := 2; n {
	// case 1, 2:
	// 	fmt.Println("1 or 2")
	// case 3, 4:
	// 	fmt.Println("3 or 4")
	// case n > 3 && n < 6:
	// 	fmt.Println("3 < n < 6")
	// default:
	// 	fmt.Println("I don't Know")
	// }

	// 比較演算子の場合
	n := 3
	switch {
	case n > 0 && n < 4:
		fmt.Println("0 < n < 4") // 0 < n < 4
	case n > 3 && n < 7:
		fmt.Println("3 < n < 7")
	default:
		fmt.Println("I don't Know")
	}
}
