package main

import (
	"fmt"
	"os"
)

// 関数が終了時に登録したdeferが実行される
func TestDefer() {
	// deferへ登録
	defer fmt.Println("END")
	fmt.Println("START")
}

// defer文を複数記述した場合は最後に登録したdefer文から呼び出される
func RunDefer() {
	defer fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3") //ここから実行される
}

func main() {
	TestDefer() // START END

	/*
		// 複数の処理をdeferへ登録する場合は無名関数を使う
		// main関数が終了した際に下記deferが呼ばれる
		defer func() {
			fmt.Println("1") // 1
			fmt.Println("2") // 2
			fmt.Println("3") // 3
		}()
	*/

	RunDefer()

	// deferによるリソース管理
	file, err := os.Create("control_struct/defer/test.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	file.Write([]byte("Hello"))
}
