package main

import (
	"fmt"
	"time"
)

func reciever(name string, ch chan int) {
	for {
		// int と　オープン状況(バッファ有無)を取得
		i, ok := <-ch
		// チャネル状態がクローズ かつ バッファが空の場合に falseになりbreakされる
		if !ok {
			break
		}
		fmt.Println(name, i)
	}
	fmt.Println(name + "END")
}

func main() {
	ch1 := make(chan int, 2)

	/*
		// closeしたチャネルには送信できない
		close(ch1)
		fmt.Println(<-ch1)
		ch1 <- 1 // panic: send on closed channel (closeしたチャネルには送信できない)
	*/

	/*
		// closeしたチャンネルから受信はできる
		close(ch1)
		fmt.Println(<-ch1) // 0
	*/

	/*
		//　2つ目の戻り値には、チャネルが閉鎖 かつ　バッファが空 の場合にfalseが返る
		close(ch1)
		i, ok := <-ch1
		fmt.Println(i, ok) // 0 false
	*/

	/*
		// チャネルがオープン または バッファに存在する　場合にtrueが返る
		ch1 <- 1
		close(ch1)
		i, ok := <-ch1
		fmt.Println(i, ok) // 1, true
	*/

	// ゴルーチンの実行
	go reciever("1.goroutin", ch1)
	go reciever("2.goroutin", ch1)
	go reciever("3.goroutin", ch1)

	i := 0
	for i < 100 {
		ch1 <- i
		i++
	}
	close(ch1)
	time.Sleep(3 * time.Millisecond)
}
