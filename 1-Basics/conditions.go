package main

import "fmt"

func main() {
	var user, pwd string
	fmt.Scan(&user, &pwd)
	if user == "user1" && pwd == "pwd1" {
		fmt.Println("Authorized!")
	} else {
		fmt.Println("Invalid Data!")
	}
}
