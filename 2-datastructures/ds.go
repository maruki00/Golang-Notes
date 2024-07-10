package main

import "fmt"

func main() {
	var x [15]int
	var twoD [2][3]int

	letters := []string{"a", "b", "c", "d"}
	letters2 := []string{}

	//Copy
	copy(letters2, letters)

	//Append Letter
	letters = append(letters, "e")

	//Delete Letter
	letters = append(letters[:1], letters[2:]...)

	var s []byte
	s = make([]byte, 5, 5)
	//OR
	ss := make([]byte, 5)

	fmt.Println(x, twoD, letters, s, ss, letters2)
}
