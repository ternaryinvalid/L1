package main

import "fmt"

type Human struct {
	age int
}

func (h Human) Walk() {
	fmt.Println("The person is walking")
}

type Action struct {
	time int
	Human
}

func main() {
	action := Action{}

	action.Walk()
}
