package main

import (
	"fmt"
	"strconv"
)

func main() {
	var s string = "100"

	// 正常
	i, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("i = %T\n", i) // i = int

	// エラーの場合
	var err_s string = "JOJO"
	i2, err := strconv.Atoi(err_s)
	if err != nil {
		fmt.Println(err) // strconv.Atoi: parsing "JOJO": invalid syntax
	}
	fmt.Printf("i2 = %T\n", i2) //i2 = int
}
