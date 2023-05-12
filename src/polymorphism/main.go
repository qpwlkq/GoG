package main

import "fmt"

type Animal interface {
	Sleep()
	GetColer() string
	GetType() string
}

type Cat struct {
	color string
}

func (this *Cat) Sleep() {
	fmt.Print("Cat is sleeping\n")
}

func (this *Cat) GetColer() string {
	return this.color
}

func (this *Cat) GetType() string {
	return "Cat"
}

type Dog struct {
	color string
}
func (this *Dog) Sleep() {
	fmt.Print("Dog is sleeping\n")
}

func (this *Dog) GetColer() string {
	return this.color
}

func (this *Dog) GetType() string {
	return "Dog"
}

func (this *Dog) GetColor() string {
	return this.color
}

func main() {
	var animal Animal
	animal = &Cat{color: "yellow"}
	animal.Sleep()
	animal = &Dog{color: "black"} 
	animal.Sleep()
	dog := Dog{color: "white"}
	fmt.Println(dog.GetColor());
}