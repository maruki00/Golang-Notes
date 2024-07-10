package main

import "fmt"

func main() {
	var age int
	var name string

	fmt.Print("Enter Age :")
	fmt.Scan(&age)
	fmt.Print("Enter Your Name : ")
	fmt.Scan(&name)

	fmt.Printf("You are %s, and you have %d years ols.\n", name, age)
}
