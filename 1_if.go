package main

import "fmt"

func main() {
	var ownMoney uint = 20000
	var priceChicken uint = 18000
	var pricePizza uint = 18000

	if priceChicken > ownMoney {
		fmt.Println("We can't buy chicken")
	} else if pricePizza > ownMoney {
		fmt.Println("We can't buy pizza")
	} else {
		fmt.Println("We can buy anything")
	}
}
