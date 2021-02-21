package main

import "fmt"

func main() {
	sl := []string{"A", "B", "C"}
	fmt.Println(sl) // [A B C]

	for index, value := range sl {
		fmt.Println(index) // 0 1 2
		fmt.Println(value) // A B C
	}

	// valueのみ使う
	for _, value := range sl {
		fmt.Println(value) // A B C
	}

	// 古典的forを使う
	for i := 0; i < len(sl); i++ {
		fmt.Println(sl[i]) // A B C
	}
}
