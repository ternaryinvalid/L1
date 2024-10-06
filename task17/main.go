package main

import "fmt"

func BinarySearch(arr []int, val int) int {
	left, right := 0, len(arr)
	mid := (left + right) / 2

	for left < right {
		if arr[mid] == val {
			return mid
		} else if arr[mid] < val {
			left = mid
		} else {
			right = mid
		}
		mid = (left + right) / 2
	}
	return -1
}

func main() {

	var sortedArray = []int{1, 2, 4, 8, 16, 25, 27, 38, 65, 77, 81}
	var searchNumber = 25
	fmt.Println("Search Number Index: ", BinarySearch(sortedArray, searchNumber))
}
