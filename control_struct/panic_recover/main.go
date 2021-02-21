package main

import "fmt"

func main() {
	// 処理を再開させる(recover)
	// 処理の再開はdefer文内のみ
	defer func() {
		if x := recover(); x != nil {
			fmt.Println(x) // runtime error
		}
	}()

	// 強制的に停止される
	panic("runtime error") // panic: runtime error
	fmt.Println("START")   // ここは実行されることがない
}
