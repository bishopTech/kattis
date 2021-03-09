package main

import (
	"fmt"
	"strings"
)

func main() { 
	var (
		input string 
		res int = 1 
	)

	fmt.Scanln(&input)
	var commands []string = strings.Split(input, "")

	for i := 0; i < len(commands); i++ {
		if commands[i] == "A" {
			res = a(res)
		} else if commands[i] == "B" {
			res = b(res)
		} else if commands[i] == "C" {
			res = c(res)
		}
	}

	fmt.Println(res)

}

func a(res int) int  {
	if res == 1 {
		res++ 
	} else if res == 2 {
		res-- 
	} 
	return res
}

func b(res int) int  {
	if res == 2 {
		res++ 
	} else if res == 3 {
		res-- 
	} 
	return res
}

func c(res int) int  {
	if res == 1 {
		res += 2
	} else if res == 3 {
		res -= 2 
	} 
	return res
}