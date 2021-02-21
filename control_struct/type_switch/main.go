package main

import "fmt"

func anything(a interface{}) {
	fmt.Println(a)
}

func main() {
	anything("aaa") // aaa
	anything(1)     // 1

	// 型アサーション
	var x interface{} = 1
	// 型アサーションでは第2戻り値にtrue or falseが返る
	i, isInt := x.(int)
	if !isInt {
		fmt.Println(isInt)
	}
	fmt.Println(i + 3) // 4

	f, isFloat64 := x.(float64)
	fmt.Println(f, isFloat64) // 0, false

	// 型アサーションとif文
	if x == nil {
		fmt.Println("None")
	} else if i, isInt := x.(int); isInt {
		fmt.Println(i, "x is Int") // 1 x is Int
	} else if s, isString := x.(string); isString {
		fmt.Println(s, isString)
	} else {
		fmt.Println("I don't Know")
	}

	// 型アサーションとswith文
	switch x.(type) {
	case int:
		fmt.Println("int") // int
	case string:
		fmt.Println("string")
	default:
		fmt.Println("I don't Know")
	}

	// 値を使う場合
	switch v := x.(type) {
	case bool:
		fmt.Println(v, "bool")
	case int:
		fmt.Println(v, "int") // 1, int
	case string:
		fmt.Println(v, "string")
	case float64:
		fmt.Println(v, "float64")
	default:
		fmt.Println("I don't Know")
	}
}
