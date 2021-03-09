package main

import "fmt"

func main() {
	var (
		numTimes int 
	)

	fmt.Scanln(&numTimes)

	for i := 0; i < numTimes; i++ {
		fmt.Printf("%d Abracadabra\n", i + 1)
	}
}