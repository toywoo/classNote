package main

import (
	"fmt"
	"unsafe"
)

func swap(ch1 *int, ch2 *int) {
	var temp int

	temp = *ch1
	*ch1 = *ch2
	*ch2 = temp
}

func main() {
	var num int = 1
	var pnum *int = &num

	fmt.Println(num, *pnum)

	fmt.Println(unsafe.Sizeof(pnum)) // 포인터의 크기는 8바이트이며 이론적으로 2^64까지 주소 접근이 가능합니다.

	var str1 string = "Hello"
	var pstr1 *string = &str1

	fmt.Println(unsafe.Sizeof(pstr1))

	//

	var num1 int = 1
	var num2 int = 2

	fmt.Println("Before: ", num1, num2)

	swap(&num1, &num2)

	fmt.Println("After: ", num1, num2)
}
