package main

import "fmt"

func MakeSequence(sequence []string) map[string]bool {
	m := make(map[string]bool)

	for _, v := range sequence {
		m[v] = true
	}

	return m
}

func main() {
	sequence := []string{"cat", "cat", "dog", "cat", "tree"}
	fmt.Println(MakeSequence(sequence))
}
