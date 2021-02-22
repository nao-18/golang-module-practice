package main

import "fmt"

type Point struct {
	A int
	B string
}

func (p *Point) String() string {
	return fmt.Sprintf("<<%v, %v>>", p.A, p.B)
}

func main() {
	p := &Point{100, "ABC"}
	fmt.Println(p) // &{100 ABC}  Stringメソッドによって書き換えられた場合 <<100, ABC>>
}
