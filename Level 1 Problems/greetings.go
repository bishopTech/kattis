package main

import "fmt"

func main() {
	var imp string
	fmt.Scanln(&imp)
	var size int = len(imp) - 2
	var out string = "h"
	for i := 0; i < size * 2; i++ {
		out += "e"
	}
	out += "y"
	fmt.Println(out)

}