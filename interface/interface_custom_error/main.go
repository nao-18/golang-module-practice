package main

import "fmt"

// カスタムエラーを定義する
type MyError struct {
	Message   string
	ErrorCode int
}

// カスタムエラーのメソッドを定義する
func (e *MyError) Error() string {
	return e.Message
}

// エラーを返す関数を定義する
// 返り値をerrorタイプとする
func RaiseError() error {
	// ポインタを生成する
	return &MyError{Message: "カスタムエラーが発生しました", ErrorCode: 1234}
}

func main() {
	err := RaiseError()
	// カスタムエラーのメソッドを表示する
	// この場合はカスタムエラー内のフィールドには直接アクセスできない
	fmt.Println(err.Error())

	// カスタムエラーのフィールドにアクセスする場合は型アサーションで実体を復元する
	e, ok := err.(*MyError)
	if ok {
		fmt.Println(e.ErrorCode)
	}

}
