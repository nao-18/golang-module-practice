package main

import "fmt"

func main() {
	var t, f bool = true, false
	fmt.Println(t, f)           // true false
	fmt.Printf("%T %T\n", t, f) // bool bool
}
