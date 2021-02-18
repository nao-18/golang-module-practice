package main

import "fmt"

func main() {
	/*
		// 型変換
		var i int = 1
		fl64 := float64(i)
		fmt.Println(fl64)               // 1
		fmt.Printf("i = %T\n", i)       // int
		fmt.Printf("fl64 = %T\n", fl64) // float64

		i2 := int(fl64)
		fmt.Printf("i2= %T\n", i2) // int

		// 注意: 浮動小数は小数点以下が切り捨てられる
		fl := 10.8
		i3 := int(fl)
		fmt.Printf("i3 = %T\n", i3) // int
		fmt.Println(i3)             // 10
	*/

	/*
		// 文字列から数値への変換
		var s string = "100"
		fmt.Printf("s = %T\n", s) // string

		// i, _ := strconv.Atoi(s) errを破棄する場合
		i, err := strconv.Atoi(s)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("i = %T\n", i) // int
		fmt.Println(i)            // 100

		// 数値から文字列への変換
		var i2 int = 200
		s2 := strconv.Itoa(i2)
		fmt.Printf("s2 = %T\n", s2) // string
		fmt.Println(s2)             // 200
	*/

	// 文字列からバイト配列への変換
	var h string = "Hello World!"
	b := []byte(h)
	fmt.Println(b)

	// バイト配列から文字列への変換
	h2 := string(b)
	fmt.Println(h2)
}
