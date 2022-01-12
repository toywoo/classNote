package main

import "fmt"

func main() {
	var result int = 0

	for i := 0; i < 1001; i++ {
		result += i
	}

	fmt.Println(result)
}
