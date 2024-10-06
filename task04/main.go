package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	ch := make(chan int)
	var nWorkers int

	fmt.Print("Введите количество воркеров: ")
	fmt.Scan(&nWorkers)

	info := make(chan os.Signal, 1)
	signal.Notify(info, syscall.SIGINT, os.Interrupt)

	var wg sync.WaitGroup

	for i := 0; i < nWorkers; i++ {
		wg.Add(1)
		go worker(ch, &wg, i)
	}

	go func() {
		for {
			select {
			case <-info:
				close(ch)
				close(info)
				return
			default:
				ch <- rand.Intn(100)
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()

	wg.Wait()
}

func worker(ch <-chan int, wg *sync.WaitGroup, id int) {
	defer wg.Done()

	for v := range ch {
		fmt.Printf("Полученое число %d из воркера %d\n", v, id+1)
	}
}
