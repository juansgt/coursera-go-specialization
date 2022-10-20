package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	var personMap map[string]string
	var name string
	var address string

	personMap = make(map[string]string)

	fmt.Println("Please enter a name:")
	fmt.Scan(&name)
	fmt.Println("Please enter an address:")
	fmt.Scan(&address)

	personMap["name"] = name
	personMap["address"] = address
	personaJson, _ := json.Marshal(personMap)
	println(string(string(personaJson)))
}
