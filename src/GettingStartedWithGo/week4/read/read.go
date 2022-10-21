package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Person struct {
	fName string
	lName string
}

func main() {
	var fileName string
	var file *os.File
	var err error
	var person *Person
	var personSplit []string
	var peopleCollection []Person

	fmt.Println("Please enter a file name:")
	fmt.Scan(&fileName)

	file, err = os.Open(fileName)

	if err == nil {
		scanner := bufio.NewScanner(file)
		scanner.Split(bufio.ScanLines)
		peopleCollection = make([]Person, 0, 5)

		for scanner.Scan() {
			person = new(Person)
			personSplit = strings.Split(scanner.Text(), " ")
			person.fName = personSplit[0]
			person.lName = personSplit[1]
			peopleCollection = append(peopleCollection, *person)
		}
		file.Close()

		for _, p := range peopleCollection {
			fmt.Println(p.fName + " " + p.lName)
		}
	} else {
		fmt.Println("An error occured trying to open the file")
	}
}
