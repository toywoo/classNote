package main

import "fmt"

func main() {
	var result int = 0
	var i int = 1

	for i < 1001 {
		result += i
		i++
	}

	fmt.Println(result)
}
