package main

import "fmt"

func main() {
	var x [15]int
	var twoD [2][3]int

	letters := []string{"a", "b", "c", "d"}

	/*  using make -> make([]T, len, cap) */
	var s []byte
	s = make([]byte, 5, 5)
	//OR
	ss := make([]byte, 5)

	fmt.Println(x, twoD, letters, s, ss)
}
