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

func mergeSlices(a []int, b []int) []int {
	final := []int{}
	i := 0
	j := 0
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			final = append(final, a[i])
			i++
		} else {
			final = append(final, b[j])
			j++
		}
	}
	for ; i < len(a); i++ {
		final = append(final, a[i])
	}
	for ; j < len(b); j++ {
		final = append(final, b[j])
	}
	return final
}

type ISortAsync interface {
	SortAsync(waitGroup *sync.WaitGroup, slice []int)
}

type IMerge interface {
	Merge(slice1 []int, slice2 []int) []int
}

type BubbleSort struct{}

func (BubbleSort *BubbleSort) SortAsync(waitGroup *sync.WaitGroup, slice []int) {
	bubbleSort(slice)
	fmt.Println("Subarray sorted in goroutine: ", slice)

	waitGroup.Done()
}

type Merge struct{}

func (Merge *Merge) Merge(slice1 []int, slice2 []int) []int {
	return mergeSlices(slice1, slice2)
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
	var merge IMerge = new(Merge)
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

	go sortAsync.SortAsync(&waitGroup, numbers1)
	go sortAsync.SortAsync(&waitGroup, numbers2)
	go sortAsync.SortAsync(&waitGroup, numbers3)
	go sortAsync.SortAsync(&waitGroup, numbers4)

	waitGroup.Wait()

	numbers = merge.Merge(merge.Merge(numbers1, numbers2), merge.Merge(numbers3, numbers4))
	fmt.Println(numbers)
}
