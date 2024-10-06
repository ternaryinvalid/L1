package main

import (
	"fmt"
	"reflect"
)

func main() {

	var inter interface{}

	inter = 2
	fmt.Println(reflect.TypeOf(inter))

	inter = "var"
	fmt.Println(reflect.TypeOf(inter))

	inter = make(chan int)
	fmt.Println(reflect.TypeOf(inter))

	inter = true
	fmt.Println(reflect.TypeOf(inter))
}
