package main

import (
	"fmt"
	"time"
)

func reciever(c chan int) {
	for {
		// channelを受け取り、変数iへ代入する
		i := <-c
		// 標準出力
		fmt.Println(i)
	}
}

func main() {
	// 双方向のchannelを作成
	ch1 := make(chan int)
	ch2 := make(chan int)

	// go routinで実行
	go reciever(ch1)
	go reciever(ch2)

	i := 0
	for i < 100 {
		// ch1, ch2へデータ送信
		ch1 <- i
		ch2 <- i
		time.Sleep(50 * time.Millisecond)
		i++
	}
}
