package main

import "fmt"

func main() {
	var num1 int = 10
	var num2 uint = 10

	fmt.Println(num1 + int(num2))

	var num3 float32 = 3.14

	fmt.Println(num1 + int(num3))
}
