package main

import (
	"fmt"
	"sync"
)

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

type ISortAsync interface {
	SortAsync(waitGroup *sync.WaitGroup, slice []int)
}

type ISort interface {
	Sort(slice []int)
}

type BubbleSort struct{}

func (BubbleSort *BubbleSort) SortAsync(waitGroup *sync.WaitGroup, slice []int) {
	bubbleSort(slice)

	waitGroup.Done()
}

func (BubbleSort *BubbleSort) Sort(slice []int) {
	bubbleSort(slice)
}

func main() {
	var number int
	var numbers []int
	var numbers1 []int
	var numbers2 []int
	var numbers3 []int
	var numbers4 []int
	var waitGroup sync.WaitGroup
	var bubbleSort = new(BubbleSort)
	var sort ISort = bubbleSort
	var sortAsync ISortAsync = bubbleSort

	waitGroup.Add(4)
	numbers = make([]int, 0, 12)

	fmt.Println("Please insert an integer an press enter (up to twelve): ")

	for i := 0; i < 12; i++ {
		fmt.Scan(&number)
		numbers = append(numbers, number)
	}

	numbers1 = numbers[0:3]
	numbers2 = numbers[3:6]
	numbers3 = numbers[6:9]
	numbers4 = numbers[9:]

	fmt.Print("Subarray to be sorted in goroutine 1: ")
	fmt.Println(numbers1)
	fmt.Print("Subarray to be sorted in goroutine 2: ")
	fmt.Println(numbers2)
	fmt.Print("Subarray to be sorted in goroutine 3: ")
	fmt.Println(numbers3)
	fmt.Print("Subarray to be sorted in goroutine 4: ")
	fmt.Println(numbers4)

	go sortAsync.SortAsync(&waitGroup, numbers1)
	go sortAsync.SortAsync(&waitGroup, numbers2)
	go sortAsync.SortAsync(&waitGroup, numbers3)
	go sortAsync.SortAsync(&waitGroup, numbers4)

	waitGroup.Wait()

	sort.Sort(numbers)
	fmt.Println(numbers)
}
