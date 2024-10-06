package main

import "fmt"

func setOne(n, i int64) int64 {
	return n | (1 << (i - 1))
}

func setZero(n, i int64) int64 {
	return n &^ (1 << (i - 1))
}

func main() {
	a := int64(20)
	fmt.Println(a)
	a = setOne(a, 2)
	fmt.Println(a)
	a = setZero(a, 2)
	fmt.Println(a)
}
