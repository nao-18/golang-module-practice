package main

import "fmt"

func Double(i int) {
	i = i * 2
}

// ポインタを引数で受け取る場合
func Doublev2(i *int) {
	*i = *i * 2
}

// 参照渡し
func Doublev3(sl []int) {
	for i, v := range sl {
		sl[i] = v * 2
	}
}

func main() {
	// 基本型は値渡しになる
	var i int = 100
	fmt.Println(i) // 100

	fmt.Println(&i) // ポインタ 0xc000016068
	Double(i)
	fmt.Println(i)

	// ポインタ型の宣言
	var p *int = &i
	fmt.Println(p) // ポインタ 0xc000016068
	// デリファレンス
	fmt.Println(*p) // 実体 100

	/*
		*p = 300        // 値の書き換え
		fmt.Println(i)  // 300
		i = 200
		fmt.Println(*p) // 200
	*/

	Doublev2(&i)   // iのポインタを渡す
	fmt.Println(i) // 200

	Doublev2(p)     // ポインタ変数を渡す
	fmt.Println(*p) // 実体 400

	// 参照型は参照渡しとなるため元の値も書き換わる
	var sl []int = []int{1, 2, 3}
	Doublev3(sl)
	fmt.Println(sl) // [2 4 6]
}
