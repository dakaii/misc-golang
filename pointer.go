package main

import "fmt"

type Dog struct {
	Name string
}

func (d Dog) Say() {
	fmt.Println("Woof!" + d.Name)
}

func Modify(d **Dog) {
	*d = &Dog{Name: "Pochi"}
}

func main() {
	d := &Dog{Name: "Bamf"}
	fmt.Println(d.Name)
	fmt.Println(&d)
	fmt.Println(&d.Name)
	d.Say()
	Modify(&d)
	d.Say()
}
