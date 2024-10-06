package main

import (
	"fmt"
	"sync"
)

func main() {
	var mass [5]int = [5]int{2, 4, 6, 8, 10}
	var wg sync.WaitGroup
	var sum int
	for _, v := range mass {
		wg.Add(1)
		go func(v int) {
			defer wg.Done()
			sum += v * v
		}(v)
	}

	wg.Wait()

	fmt.Println(sum)
}
