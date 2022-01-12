package main

import "fmt"

func main() {
	var logi1 bool = true
	var logi2 bool = false
	var result1 bool

	result1 = logi1 || logi2

	fmt.Println(result1)

	logi1 = false
	logi2 = false

	result1 = logi1 || logi2

	fmt.Println(result1)
	fmt.Println(!result1) // NOT
}
