package main

import "fmt"

func main() {
	var imp int
	fmt.Scanln(&imp)
	if (imp % 2 == 0) {
		fmt.Println("Bob")
	} else {
		fmt.Println("Alice")
	}

}