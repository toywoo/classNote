package main

import "fmt"

func main() {
	var num1 uint = 18446744073709551615
	var num2 uint = 1

	fmt.Println(num1 << num2)
}
