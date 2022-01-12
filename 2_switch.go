package main

import "fmt"

func main() {
	var wantedFood string = "C"

	switch wantedFood {
	case "C":
		fmt.Println("He want chicken")
		break
	case "P":
		fmt.Println("He want Pizza")
		fallthrough
	default:
		fmt.Println("I don't know wanted food")
	}
}
