package main

import "fmt"

func main() {
	var (
		limit int 
		numInput int
		temp int  
		listOfUsage []int
	)

	fmt.Scan(&limit)
	fmt.Scan(&numInput)
	for i := 0; i < numInput; i++ {
		fmt.Scan(&temp)
		listOfUsage = append(listOfUsage, temp)
	}

	fmt.Println(calcBalance(limit, listOfUsage))

}

func calcBalance(limit int, list []int ) int {
	var res int = 0
	for i:= 0; i < len(list); i++ {
		res += limit - list[i]
	}
	return res + limit
}