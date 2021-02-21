package main

import "fmt"

func main() {
	// 宣言(整数型のスライス)
	var sl []int
	fmt.Println(sl) // []

	// 明示的宣言
	var sl2 []int = []int{1, 3, 4}
	fmt.Println(sl2) // [1 3 4]

	// 暗黙的宣言
	sl3 := []string{"A", "B"}
	fmt.Println(sl3) // [A B]

	// make関数を使ってスライスを生成する場合
	// 第二引数に要素数を指定する
	sl4 := make([]int, 6)
	fmt.Println(sl4) // [0 0 0 0 0 0]

	// 値の代入
	sl2[0] = 1000
	fmt.Println(sl2) // [1000 3 4]

	// 値の取り出し
	sl5 := []int{1, 2, 3, 4, 5}
	fmt.Println(sl5[0]) // 1
	// インデックス番号(2~4)を取り出す
	fmt.Println(sl5[2:4]) // [3 4]
	// インデックス番号(2)より前全てを取り出す
	fmt.Println(sl5[:2]) // [1 2]
	// インデックス番号(2)以降全てを取り出す
	fmt.Println(sl5[2:]) // [3 4 5]
	// 配列の全てを取り出す
	fmt.Println(sl5[:]) // [1 2 3 4 5]
	// 最後のみ取得する
	fmt.Println(len(sl5) - 1) // 4
	// 最初と最後以外すべてを取得する
	fmt.Println(sl5[1 : len(sl5)-1])
}
