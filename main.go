package main

import "fmt"

func main() {
	var arr = [5]int{1, 2, 3, 4, 5}
	var sli = []int{1, 2, 3, 4, 5}
	var num int

	num = arr[0]
	num = sli[0]

	fmt.Println(num)

	//

	sli = append(sli, 6)
	fmt.Println(sli)

	//

	for index, value := range arr {
		fmt.Println("\n", index, value)
	}

	//

	const (
		KOREAN = iota
		MATH
		ENGLISH
	)

	var mulDiArr = [3][5]int{
		{1, 2, 3, 4, 5},
		{1, 2, 3, 4, 5},
		{1, 2, 3, 4, 5},
	}

	fmt.Println(mulDiArr[KOREAN][2])

	//

	var gradeMap = make(map[int][]int, 3)

	gradeMap[KOREAN] = []int{1, 2, 3, 4, 5}
	gradeMap[MATH] = []int{1, 2, 3, 4, 5}
	gradeMap[ENGLISH] = []int{1, 2, 3, 4, 5}

	fmt.Println(gradeMap[KOREAN][2])
}
