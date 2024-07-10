package main

import "fmt"

func main() {
	var val int
	fmt.Scan(&val)
	switch val {
	case 1:
		fmt.Println("You entered 1")
		return
	case 2:
		fmt.Println("You entered 2")
		return
	default:
		fmt.Println("Invalid Choise")
		return
	}
}
