package main

import "fmt"

func main() {
	var logi1 bool = false
	var logi2 bool = false
	var result bool

	result = logi1 == logi2

	fmt.Println(result)

	var num1 int = 3
	var num2 int = 6

	result = num1 == num2

	fmt.Println(result)

	result = num1 != num2

	fmt.Println(result)

	result = num1 != num2

	fmt.Println(!result) // NOT

	var str1 string = "Hello"
	var str2 string = "Hello "

	result = str1 == str2

	fmt.Println(result)
}
