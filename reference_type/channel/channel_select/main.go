package main

import "fmt"

func main() {
	ch1 := make(chan int, 2)
	ch2 := make(chan string, 2)

	/*
		ch2 <- "A"

		v1 := <-ch1 // ここでエラーになる
		v2 := <-ch2
		fmt.Println(v1)
		fmt.Println(v2)
	*/

	/*
		// selectはcaseがランダムに実行される
		ch2 <- "A"
		select {
		case v1 := <-ch1:
			fmt.Println(v1 + 1000)
		case v2 := <-ch2:
			fmt.Println(v2 + "!?")
		}
	*/

	// selectのdefault
	/*
		ch2 <- "A"
		ch1 <- 1
		ch2 <- "B"
		ch1 <- 2
	*/

	select {
	case v1 := <-ch1:
		fmt.Println(v1 + 1000)
	case v2 := <-ch2:
		fmt.Println(v2 + "!?")
	default:
		fmt.Println("どちらでもない")
	}

	ch3 := make(chan int)
	ch4 := make(chan int)
	ch5 := make(chan int)

	//reciever
	go func() {
		for {
			// 変数iへch3を受信
			i := <-ch3
			// ch4へ変数iを倍にした値を送信
			ch4 <- i * 2
		}
	}()

	go func() {
		for {
			// 変数i2へch4を受信
			i2 := <-ch4
			// ch5へ変数i2から-1した値を送信
			ch5 <- i2 - 1
		}
	}()

	n := 0
	for {
		select {
		case ch3 <- n:
			n++
		case i3 := <-ch5:
			fmt.Println("reciever", i3)
		}
		if n > 100 {
			break
		}
	}

	// defultを使う場合
Loop:
	for {
		select {
		case ch3 <- n:
			n++
		case i3 := <-ch5:
			fmt.Println("reciever", i3)
		default:
			if n > 100 {
				break Loop
			}
		}
	}
}
