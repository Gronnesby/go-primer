package main

import (
	"fmt"
	"sync"
)

// START OMIT
func main() {

	go parallelWork(1, "Hello")
}

func parallelWork(idx int, msg string) {
	fmt.Println(idx, msg)
}

// FIRST OMIT
func main() {
	var wg sync.WaitGroup

	for i := 0; i < 40; i++ {
		wg.Add(1)

		go func(idx int, msg string) {
			defer wg.Done()
			fmt.Println(idx, msg)
		}(i, "Hello")
	}
	wg.Wait()
}

// LAST OMIT
