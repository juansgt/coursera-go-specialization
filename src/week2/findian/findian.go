package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var word string
	var resultMessage string
	var count int
	var scanner bufio.Scanner

	scanner = *bufio.NewScanner(os.Stdin)

	fmt.Println("Please enter something:")
	scanner.Scan()
	word = scanner.Text()

	count = 0
	for i, letter := range word {
		switch i {
		case 0:
			count = increaseCountIfEquals(string(letter), "i", count)
		case len(word) - 1:
			count = increaseCountIfEquals(string(letter), "n", count)
		default:
			count = increaseCountIfEquals(string(letter), "a", count)
		}
	}

	if count == 3 {
		resultMessage = "Found!"
	} else {
		resultMessage = "Not Found!"
	}

	fmt.Println(resultMessage)
}

func increaseCountIfEquals(value string, target string, count int) int {
	if strings.ToLower(value) == target {
		count = count + 1
	}

	return count
}
