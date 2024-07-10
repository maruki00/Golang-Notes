package main

import "fmt"

func main() {
	var intPointer *int
	var intval int = 10
	intPointer = &intval

	fmt.Println(intPointer, *intPointer, intval)
}
