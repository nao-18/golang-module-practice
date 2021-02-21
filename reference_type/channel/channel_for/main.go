package main

import "fmt"

func main() {
	ch1 := make(chan int, 3)
	ch1 <- 1
	ch1 <- 2
	ch1 <- 3
	// チャネルのforはclose()して使うのが基本
	// オープン状態だとfatal error: all goroutines are asleep - deadlock!になる
	close(ch1)
	for i := range ch1 {
		fmt.Println(i)
	}
}
