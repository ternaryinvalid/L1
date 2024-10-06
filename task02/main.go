package main

import (
	"fmt"
	"sync"
)

func main() {
	var mass [5]int = [5]int{2, 4, 6, 8, 10}
	var wg sync.WaitGroup

	for _, v := range mass {
		wg.Add(1)

		go func(v int) {
			defer wg.Done()

			fmt.Println(v * v)
		}(v)
	}

	wg.Wait()

}
