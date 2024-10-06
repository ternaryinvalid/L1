package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// Закрытие через канал
func MethodChannel() {
	ch := make(chan int)

	go func() {

		for {
			select {
			case <-ch:
				close(ch)
				fmt.Println("Горутина на каналах завершила работу")
				return
			default:
				fmt.Println("Горутина на каналах")
				time.Sleep(time.Second)
			}
		}
	}()

	time.Sleep(3 * time.Second)
	ch <- 1 // посылаем сигнал для завершения горутины

}

// Закрытие через WaitGroup
func MethodWaitGroup() {
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		fmt.Println("Горутина с WaitGroup работает")
	}()

	wg.Wait()
	fmt.Println("Горутина с WaitGroup завершила работу")
}

// Завершение с помощью контекста
func MethodContext() {
	ctx, cancel := context.WithCancel(context.Background())
	defer fmt.Println("Горутина с контекстом завершила работу")
	defer cancel()

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				fmt.Println("Горутина с контекстом работает")
				time.Sleep(time.Second)
			}
		}
	}()

	time.Sleep(2 * time.Second)

}

func main() {
	MethodChannel()
	fmt.Println()

	MethodWaitGroup()
	fmt.Println()

	MethodContext()
	fmt.Println()

	fmt.Println("Все Горутины завершили свои работы, программа завершена")
}
