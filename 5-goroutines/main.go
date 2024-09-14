package main

import (
	"fmt"
	"sync"
)

func CountSum(mx int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := range 10 {
		fmt.Println(i + mx)
	}
}

func main() {
	wg := &sync.WaitGroup{}
	for i := range 10 {
		wg.Add(1)
		go CountSum(i*10, wg)
	}
	wg.Wait()
}
