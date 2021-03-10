package main

import (
	"fmt"
	"math"
)

func main() {
	var (
		numOfInput int 
		temp float64 
		exponent float64
		base float64
		res int = 0 
	)

	fmt.Scan(&numOfInput)
	for i := 0; i < numOfInput; i++ { 
		fmt.Scan(&temp)
		exponent = math.Remainder(temp, 10)
		base = math.Floor(temp/10)
		res += int(math.Pow(base,exponent))
	}
	fmt.Println(res)
}


