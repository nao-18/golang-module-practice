package main

import "fmt"

func Later() func(string) string {
	var store string
	return func(next string) string {
		s := store
		store = next
		return s
	}
}

func main() {
	f := Later()
	fmt.Println(f("Hello"))  //
	fmt.Println(f("My"))     // Hello
	fmt.Println(f("name"))   // My
	fmt.Println(f("is"))     // name
	fmt.Println(f("Golang")) // is
}
