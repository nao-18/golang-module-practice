package main

/*
	アルファベット順でインポートが推奨
*/
import (
	// 別名を任意でつけられる
	f "fmt"

	"example.com/sample-app/scope/foo"
)

func main() {
	f.Println(foo.Max) // 100
	// 一文字目が小文字のため、プライベートになり参照できない
	// fmt.Println(foo.min)

	// ゲッターで呼び出せる
	f.Println(foo.ReturnMin()) // 1
}
