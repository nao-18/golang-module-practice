package main

import "fmt"

// func Plus(x int, y int) int {
// 	return x + y
// }

//型が同じ場合は型を省くことができる
func Plus(x, y int) int {
	return x + y
}

//複数の返り血がある場合
func Div(x, y int) (int, int) {
	q := x / y
	r := x % y
	return q, r
}

// 返り値に変数名を指定すると、returnに描かなくても良い
func Double(price int) (result int) {
	result = price * 2
	return
}

// 返り値がない関数
func Noreturn() {
	fmt.Println("No Return")
	return
}

func main() {
	i := Plus(1, 4)
	fmt.Println(i) // 4

	i2, i3 := Div(9, 1)
	fmt.Println(i2, i3) // 9 0

	i4 := Double(1000)
	fmt.Println(i4) // 2000

	Noreturn() // No Return
}
