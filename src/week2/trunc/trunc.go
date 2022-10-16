package main

import "fmt"

func main() {
	var entryFloatNumber float32

	fmt.Println("Please enter a floating point number:")
	fmt.Scan(&entryFloatNumber)
	fmt.Println(int32(entryFloatNumber))
}
