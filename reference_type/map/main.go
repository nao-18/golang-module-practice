package main

import "fmt"

func main() {
	// 明示的宣言 map[key-type]value-type
	var m = map[string]int{"apple": 100, "orange": 200}
	fmt.Println(m) // map[apple:100 orange:200]

	// 暗黙的宣言
	m2 := map[string]int{"apple": 200, "orange": 300}
	fmt.Println(m2) // map[apple:200 orange:300]

	// 改行した場合
	m3 := map[int]string{
		1: "A",
		2: "B",
	}
	fmt.Println(m3) // map[1:A 2:B]

	// make関数でmapを生成した場合
	m4 := make(map[int]string)
	fmt.Println(m4) // map[]
	// 値を代入
	m4[1] = "Japan"
	m4[2] = "Usa"
	fmt.Println(m4) // map[1:Japan 2:Usa]

	// 値の取り出し
	fmt.Println(m4[1])      // Japan
	fmt.Println(m["apple"]) // 100
	fmt.Println(m4[100])    // 存在しないkeyの場合は空文字が返る

	//　取り出しに成功したかtrue or falseが返る
	s, ok := m4[1]
	fmt.Println(s, ok) // Japan true

	// 値の追加
	m4[3] = "China"
	fmt.Println(m4) // map[1:Japan 2:Usa 3:China]

	// 要素の削除はdelete()メソッドを使う
	delete(m4, 3)
	fmt.Println(m4) // map[1:Japan 2:Usa]

	// 要素数 len()
	fmt.Println(len(m4)) // 2
}
