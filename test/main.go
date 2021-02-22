package main

import (
	"fmt"

	"example.com/sample-app/test/alib"
)

func IsOne(i int) bool {
	if i == 1 {
		return true
	} else {
		return false
	}
}

func main() {
	fmt.Println(IsOne(1)) // true
	fmt.Println(IsOne(0)) // false

	s := []int{1, 2, 3, 4, 5}
	fmt.Println(alib.Average(s)) // 3
}
