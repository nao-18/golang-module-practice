package main

import "fmt"

// パッケージの初期化はinitを使う

func init() {
	fmt.Println("init")
}

func main() {
	fmt.Println("main")
}
