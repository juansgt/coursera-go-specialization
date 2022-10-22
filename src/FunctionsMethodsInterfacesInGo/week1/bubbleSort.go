package main

import "fmt"

func bubbleSort(slice []int) {
	for i := 0; i < len(slice)-1; i++ {
		for j := 0; j < len(slice)-i-1; j++ {
			if slice[j] > slice[j+1] {
				swap(slice, j)
			}
		}
	}
}

func swap(slice []int, i int) {
	slice[i], slice[i+1] = slice[i+1], slice[i]
}

func main() {
	var number int
	var numbers []int
	numbers = make([]int, 0, 10)

	fmt.Println("Please enter ten integers:")

	for i := 0; i <= 9; i++ {
		fmt.Scan(&number)
		numbers = append(numbers, number)
	}

	bubbleSort(numbers)
	fmt.Println(numbers)
}
