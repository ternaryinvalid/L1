package main

import "fmt"

func Write(nums []int) <-chan int {
	ch := make(chan int)

	go func() {
		defer close(ch)

		for _, v := range nums {
			ch <- v
		}

	}()

	return ch
}

func PowChan(ch <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)
		for v := range ch {
			out <- v * v
		}
	}()

	return out
}

func main() {
	ch := Write([]int{0, 1, 2, 3, 4})
	chPow := PowChan(ch)

	for v := range chPow {
		fmt.Println(v)
	}

}
