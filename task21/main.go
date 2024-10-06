package main

import (
	"fmt"
	"math/big"
)

func main() {

	a := big.NewInt(9688608)
	b := big.NewInt(2453495)

	//сложение
	add := big.NewInt(0)
	add = add.Add(a, b)
	fmt.Println(add)

	//вычитание
	sub := big.NewInt(0)
	sub = sub.Sub(a, b)
	fmt.Println(sub)

	//умножение
	mul := big.NewInt(0)
	mul = mul.Mul(a, b)
	fmt.Println(mul)

	//деление
	div := big.NewInt(0)
	div = div.Div(a, b)
	fmt.Println(div)

}
