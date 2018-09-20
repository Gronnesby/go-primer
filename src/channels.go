package main

import (
	"fmt"
	"time"
)

func main() {
	p1 := make(chan bool)
	go heavyWork(p1, 5)
	<-p1
}

func heavyWork(c chan bool, work int) {
	fmt.Println("Working")
	duration := time.Duration(work * int(time.Second))
	time.Sleep(duration)
	fmt.Println("Done")

	c <- true
}
