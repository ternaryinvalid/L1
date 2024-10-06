package main

import (
	"fmt"
	"time"
)

func sleep(sec time.Duration) {
	<-time.After(sec)
}

func main() {
	fmt.Println("main fall asleep for 3 seconds")
	sleep(3 * time.Second)
	fmt.Println("the program is ended")
}
