package main

import "fmt"

func main() {
	var x [15]int
	var twoD [2][3]int

	letters := []string{"a", "b", "c", "d"}
	var letters2 []string

	//Copy
	copy(letters2, letters)

	//Append Letter
	letters = append(letters, "e")

	//Delete Letter
	letters = append(letters[:1], letters[2:]...)

	fmt.Println(x, twoD, letters, letters2)

	hashMap := make(map[int]int)
	hashMap[1] = 120
	hashMap[45] = 324

}
