package main

import "fmt"

func main() {
	var nums [3]int
	for i := 0; i < 3; i++ {
		nums[i] = (i << 1) * i
		fmt.Println("index : ", i, nums[i])
	}
}
