package main

import "fmt"

func main() {

	var name string

	name = "xishanyu"

	var allType interface{}
	allType = name

	value, _ := allType.(string)
	fmt.Println(value)

	value := 

}