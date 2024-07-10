package main

import "fmt"

func function1() {
	fmt.Println("Functin With params and return value ")
}

func function2() int {
	return 1
}


func function3(a, b int) (int, int) {
	return b, a
}

func functionWithCallback(callback func(int) int ) int {
	return callback(1)
}

func main() {
	fmt.Println("result : ", functionWithCallback(func(a int) int {
		return a
	}))
}