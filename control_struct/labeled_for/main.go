package main

import "fmt"

func main() {
	/*
			// ラベル付きfor
		Loop:
			for {
				for {
					for {
						fmt.Println("Srat")
						break Loop
					}
					fmt.Println("処理をしない")
				}
				fmt.Println("処理をしない")
			}
			fmt.Println("End")
	*/

	// ラベル付きforとcontinue
	// Loopで一番その側のforへ飛ばす
Loop:
	for i := 0; i < 3; i++ {
		for j := 1; j < 3; j++ {
			if j > 1 {
				continue Loop
			}
			fmt.Println(i, j, i*j)
		}
		fmt.Println("処理をしない")
	}
}
