package main

import "fmt"

func main() {
	var result int = 0
	var i int = 1

addOne:
	result += i
	i++

	if i < 1001 {
		goto addOne
	}

	fmt.Println(result)
}
