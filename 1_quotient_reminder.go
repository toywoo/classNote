package main

import "fmt"

func main() {
	var num1 int = 1
	var num2 int = 2
	var result1 int
	var result2 int

	result1 = num1 / num2 // 0
	result2 = num1 % num2 // 1

	fmt.Println(result1, result2)
}
