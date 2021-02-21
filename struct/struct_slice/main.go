package main

import "fmt"

type Person struct {
	Name string
}

type Persons struct {
	Persons []*Person
}

func main() {
	/*
		// 構造体Personのスライス生成
		ps := make([]Person, 5)
		fmt.Println(ps) // [{} {} {} {} {}]

		ps[0].Name = "Mike"
		ps[1].Name = "Jon"
		ps[2].Name = "Nina"
		ps[3].Name = "Joe"
		ps[4].Name = "Nancy"

		fmt.Println(ps) // [{Mike} {Jon} {Nina} {Joe} {Nancy}]
	*/

	p1 := Person{"Mike"}
	p2 := Person{"Jon"}
	p3 := Person{"Nina"}
	p4 := Person{"Jone"}
	p5 := Person{"Nancy"}

	ps := Persons{}
	ps.Persons = append(ps.Persons, &p1, &p2, &p3, &p4, &p5)

	for _, p := range ps.Persons {
		fmt.Println(p)
		// &{Mike}
		// &{Jon}
		// &{Nina}
		// &{Jone}
		// &{Nancy}
	}
}
