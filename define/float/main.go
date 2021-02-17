package main

import "fmt"

func main() {
	var fl64 float64 = 2.4
	fmt.Println(fl64)

	// 暗黙的な定義時はfloat64が初期型になる
	fl := 3.2
	fmt.Printf("%T\n", fl) //float64

	//　明示的にfloat32
	var fl32 float32 = 3.5
	fmt.Printf("%T\n", fl32) // float32

	// 型変換
	fmt.Printf("%T\n", float64(fl32)) // float64

	// 正の無限大 +Inf 負の無限大 -Inf 非数 NaN
	zero := 0.0
	pinf := 1.0 / zero
	fmt.Println(pinf)

	ninf := -1.0 / zero
	fmt.Println(ninf)

	nan := zero / zero
	fmt.Println(nan)
}
