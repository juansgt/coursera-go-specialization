package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	var element string
	var elementInt int
	var err error
	var numbers []int
	numbers = make([]int, 0, 3)

	element = "0"
	for element != "X" {
		fmt.Println("Please enter an integer number or X to exit:")
		fmt.Scan(&element)
		elementInt, err = strconv.Atoi(element)

		if err == nil {
			numbers = append(numbers, elementInt)
			sort.Slice(numbers, func(i, j int) bool { return numbers[i] < numbers[j] })
			fmt.Println("By name:", numbers)
		}
	}
}
