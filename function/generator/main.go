package main

import "fmt"

//クロージャー(ジェネレータ)
func integers() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	ints := integers()  // 定義
	fmt.Println(ints()) // 1
	fmt.Println(ints()) // 2
	fmt.Println(ints()) // 3
	fmt.Println(ints()) // 4
	fmt.Println(ints()) // 5
	fmt.Println(ints()) // 6

	otherints := integers()  // 定義
	fmt.Println(otherints()) // 1
	fmt.Println(otherints()) // 2
	fmt.Println(otherints()) // 3
	fmt.Println(otherints()) // 4
	fmt.Println(otherints()) // 5
	fmt.Println(otherints()) // 6
	fmt.Println(otherints()) // 7
	fmt.Println(otherints()) // 8
	fmt.Println(otherints()) // 9

}
