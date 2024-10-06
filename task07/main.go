package main

import (
	"fmt"
	"sync"
)

func main() {
	mapa := map[int]int{}

	wg := sync.WaitGroup{}
	mu := sync.Mutex{}

	wg.Add(1)

	go func() {
		defer wg.Done()

		for i := 0; i < 5; i++ {
			mu.Lock()
			mapa[i] = i
			mu.Unlock()
			fmt.Println(i, "is written is map concurently")
		}
	}()

	wg.Wait()
	fmt.Println(mapa)
}
