package main

import "fmt"

func reverse(str string) string {
	var result = []rune(str)
	//Метод двух указателей
	left, right := 0, len(result)-1

	for left < right {
		result[left], result[right] = result[right], result[left]
		left++
		right--
	}

	return string(result[:])
}

func main() {
	var testSting = "главрыба"
	fmt.Println(testSting, reverse(testSting))
}
