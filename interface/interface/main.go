package main

import "fmt"

// &変数　ポインタ　アドレスを取得する
// *変数 ポインタ表現 実体を取得する

type Stringfy interface {
	ToString() string
}

type Person struct {
	Name string
	Age  int
}

func (p *Person) ToString() string {
	return fmt.Sprintf("Name=%v, Age=%v", p.Name, p.Age)
}

type Car struct {
	Number string
	Model  string
}

func (c *Car) ToString() string {
	return fmt.Sprintf("Number=%v, Model=%v", c.Number, c.Model)
}

func main() {
	vs := []Stringfy{
		&Person{Name: "Taro", Age: 21},
		&Car{Number: "100-0011", Model: "AB-1234"},
	}

	for _, v := range vs {
		fmt.Println(v.ToString())
	}
}
