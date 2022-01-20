package main

import "fmt"

func main() {
	var num1 int = 1

	fmt.Println(&num1)
	fmt.Println(*&num1)
}
