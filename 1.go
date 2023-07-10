package main

import (
	"fmt"
)

type Human1 struct {
	Name string
	Age  int
}

func (h *Human1) SayHello() {
	fmt.Println("Hello, my name is", h.Name)
}

type Action struct {
	Human1
}

func (a *Action) Walk() {
	fmt.Println(a.Name, "is walking")
}

func main1() {
	a := Action{
		Human1: Human1{
			Name: "John Doe",
			Age:  30,
		},
	}

	a.SayHello() // Вызываем метод SayHello из экземпляра Action
	a.Walk()
}
