package main

import "fmt"

func main() {
	var result int = 0
	var i int = 0

	for {
		result += i
		if i > 1001 {
			fmt.Println("escape!", i)
			break
		}
	}

	fmt.Println(result)
}
