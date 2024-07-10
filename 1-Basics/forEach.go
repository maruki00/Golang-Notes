package main_for_each_loop

import "fmt"

func main() {
	arr := []int{1, 2, 4, 5, 78, 8}
	for i := range 10 {
		fmt.Println("index : ", i)
	}

	for index, item := range arr {
		fmt.Printf("index : %d, value : %d.\n")
	}
}