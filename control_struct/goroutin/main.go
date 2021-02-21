package main

import (
	"fmt"
	"time"
)

/*
 並行処理
*/

func sub() {
	for {
		fmt.Println("Sub Loop")
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	go sub() // Sub Loop
	go sub() // Sub Loop

	for {
		fmt.Println("Main Loop") // Main Loop
		time.Sleep(200 * time.Millisecond)
	}
}
