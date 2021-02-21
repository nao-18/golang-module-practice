package main

import "fmt"

/*
 チャネルは複数のゴルーチン間のデータ受け渡しに使用
*/

func main() {
	// 宣言 双方向
	var ch1 chan int

	/*
		// 受信専用チャネルの宣言
		var ch2 <-chan int

		// 送信専用チャネルの宣言
		var ch3 chan<- int
	*/

	// make関数でチャネル初期化し、読み込み書き込み可能にする
	ch1 = make(chan int)

	// 暗黙的定義
	ch2 := make(chan int)

	// バッファサイズの調べ方
	fmt.Println(cap(ch1)) // 0
	fmt.Println(cap(ch2)) // 0

	// バッファサイズを指定する場合
	ch3 := make(chan int, 5)
	fmt.Println(cap(ch3)) // 5 (バッファサイズ)

	// データの送信
	ch3 <- 1
	// データ数を調べる len
	fmt.Println(len(ch3)) // 1 (取り出した値)

	ch3 <- 2
	ch3 <- 3
	fmt.Println("len", len(ch3)) // len 3 (ch3の要素数)

	// データの受信
	// 受信は先入先出になる(キュー)
	i := <-ch3
	fmt.Println(i)               // 1 (取り出した値)
	fmt.Println("len", len(ch3)) // len 2 (ch3の要素数)
	i2 := <-ch3
	fmt.Println(i2)              // 2 (取り出した値)
	fmt.Println("len", len(ch3)) // len 1 (ch3の要素数)

	fmt.Println(<-ch3)           // 3 (取り出した値)
	fmt.Println("len", len(ch3)) // len 0 (ch3の要素数)

	ch3 <- 1
	ch3 <- 2
	ch3 <- 3
	ch3 <- 4
	ch3 <- 5
	ch3 <- 6 // fatal error: all goroutines are asleep - deadlock! (バッファサイズが5のため)
}
