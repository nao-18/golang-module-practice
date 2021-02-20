package main

import "fmt"

func main() {
	/*
		i := 0
		for {
			i++
			if i > 3 {
				break
			}
			fmt.Println(i) // 1 2 3
		}
	*/

	/*
		point := 0
		for point < 10 {
			fmt.Println(point)
			point++
		}
	*/

	/*
		for i := 0; i < 10; i++ {
			if i == 7 {
				continue
			}
			if i == 3 {
				break
			}
			fmt.Println(i) // 0 1 2
		}
	*/

	/*
		arr := [3]int{11, 3, 4}
		for i := 0; i < len(arr); i++ {
			fmt.Println(arr[i]) // 11 3 4
		}
	*/

	/*
		// golangはforeachがないため下記で配列を回す
		arr := [4]int{2, 3, 4, 5}
		for index, value := range arr {
			fmt.Println(index) // 0 1 2 3
			fmt.Println(value) // 2 3 4 5
		}
	*/

	//スライス
	sl := []string{"Python", "PHP", "Rust"}
	for index, value := range sl {
		fmt.Println(index) // 0 1 2
		fmt.Println(value) // Python PHP Rust
	}

	//Map
	m := map[string]int{"apple": 100, "orange": 300}
	for index, value := range m {
		fmt.Println(index) // apple orange
		fmt.Println(value) // 100 300
	}
}
