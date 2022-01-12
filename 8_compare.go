package main

import "fmt"

func main() {
	var num1 int = 3
	var num2 int = 6
	var result bool

	result = num1 <= num2

	fmt.Println(result)

	result = num1 >= num2

	fmt.Println(result)

	num1 = 3
	num2 = 3

	result = num1 > num2

	fmt.Println(result)
}
