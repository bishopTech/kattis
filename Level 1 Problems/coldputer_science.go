package main

import (
	"fmt"
)

func main() {
	var (
		numIn int 
		intArr []int
		temp int
	)

	fmt.Scanln(&numIn)

	for i := 0; i < numIn; i++ {
		fmt.Scan(&temp)
		if temp < 0 {
			intArr = append(intArr, temp)
		}
	}

	fmt.Println(len(intArr))

}