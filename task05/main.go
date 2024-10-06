package main

import (
	"fmt"
	"time"
)

const N = 5

func sender(ch chan<- int) {
	for i := 0; ; i++ {
		ch <- i
		time.Sleep(500 * time.Millisecond)
	}
}

func receiver(ch <-chan int) {
	for val := range ch {
		fmt.Println(val)
	}
}

func main() {
	ch := make(chan int)
	timeoutChan := time.After(N * time.Second)

	go sender(ch)

	go receiver(ch)

	<-timeoutChan

	fmt.Println("Конец программы")

}
