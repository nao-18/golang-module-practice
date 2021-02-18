package main

import "fmt"

/*
	配列は後から追加、変更等ができない
*/

func main() {
	// 明示的な定義
	var arr1 [3]int
	fmt.Println(arr1)        // [0 0 0]
	fmt.Printf("%T\n", arr1) // [3]int [3int]が型になる

	// 値を持たせた場合
	var arr2 [3]string = [3]string{"A", "B", "C"}
	fmt.Println(arr2)        // [A B C]
	fmt.Printf("%T\n", arr2) // [3]string

	// 暗黙的な定義
	arr3 := [3]int{1, 2}
	fmt.Println(arr3)        // [1 2 0]
	fmt.Printf("%T\n", arr3) // [3]int

	// 要素数を明示しない場合は[...]を使う
	arr4 := [...]string{"C", "D"}
	fmt.Println(arr4)        // [C D]
	fmt.Printf("%T\n", arr4) // [2]string

	// 値の取り出し
	fmt.Println(arr1[0]) // 0
	// 最後の要素の取り出し [要素数-1]
	fmt.Println(arr2[3-1]) // C

	// 値の更新
	arr1[0] = 1
	fmt.Println(arr1[0]) // 1

	// 要素数が違う場合は代入できない
	// var arr5 [4]int
	// arr5 = arr1
	// fmt.Println(arr5)

	// 要素数の確認
	fmt.Println(len(arr1)) // 3
}
