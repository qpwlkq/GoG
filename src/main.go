package main

import "fmt"

type Hero struct {
	Name string
	Ad   int
	Level int
}

type Person struct {
	Hero
	level int
}

func (this *Person) SetName3(name string) {
	this.Name = name;
}

func (this Hero) GetName() {
	fmt.Println(this.Name);
}

func (this Hero) SetName(name string) {
	this.Name = name;
}

func (this *Hero) SetName2(name string) {
	this.Name = name;
}

func main() {
	hero := Hero{Name: "张三", Ad: 100, Level: 0}
	hero.SetName("李四")
	person := Person{Hero: hero, level: 1}
	person.SetName3("王五")
	person.GetName()
	hero.GetName()
}