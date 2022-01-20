package main

import (
	"fmt"
)

func add(op1 int, op2 int) int {
	return op1 + op2
}

func minus(op1 int, op2 int) int {
	return op1 - op2
}

func main() {
	var num1 int = 1
	var num2 int = 2
	var result int = add(num1, num2) + minus(num1, num2)

	fmt.Println(result)
}
