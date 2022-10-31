package main

import (
	"fmt"
)

type Animal interface {
	Eat() string
	Move() string
	Speak() string
}

type Cow struct{}
type Bird struct{}
type Snake struct{}

// Cow

func (cow *Cow) Eat() string {
	return "grass"
}

func (cow *Cow) Move() string {
	return "walk"
}

func (cow *Cow) Speak() string {
	return "moo"
}

// Bird

func (bird *Bird) Eat() string {
	return "worms"
}

func (bird *Bird) Move() string {
	return "fly"
}

func (bird *Bird) Speak() string {
	return "peep"
}

// Snake

func (snake *Snake) Eat() string {
	return "mice"
}

func (snake *Snake) Move() string {
	return "slither"
}

func (snake *Snake) Speak() string {
	return "hsss"
}

func main() {
	var animalMap map[string]Animal
	var command string
	var animalName string
	var animalType string
	var animalInformation string

	animalMap = make(map[string]Animal)

commandLoop:
	for {
		fmt.Print("> ")
		fmt.Scan(&command)

		switch command {
		case "newanimal":
			fmt.Scan(&animalName)
			fmt.Scan(&animalType)
			createAnimal(animalName, animalType, animalMap)
		case "query":
			fmt.Scan(&animalName)
			fmt.Scan(&animalInformation)
			fmt.Println(getAnimalInformation(animalName, animalInformation, animalMap))
		case "X":
			break commandLoop
		default:
			fmt.Println("Don't know that commnad")
		}
	}
}

func getAnimalInformation(name string, animalInformation string, animalMap map[string]Animal) string {
	var requestedInformation string
	animal, keyExists := animalMap[name]

	if keyExists {
		switch animalInformation {
		case "eat":
			requestedInformation = animal.Eat()
		case "move":
			requestedInformation = animal.Move()
		case "speak":
			requestedInformation = animal.Speak()
		default:
			fmt.Println("Don't know that type of information, only eat, move or speak")
		}
	} else {
		fmt.Println("Can't find information for that name")
	}

	return requestedInformation
}

func createAnimal(name string, animalType string, animalMap map[string]Animal) {
	created := true

	switch animalType {
	case "cow":
		animalMap[name] = new(Cow)
	case "bird":
		animalMap[name] = new(Bird)
	case "snake":
		animalMap[name] = new(Snake)
	default:
		created = false
		fmt.Println("Don't know that type of animal")
	}

	if created {
		fmt.Println("Created it!")
	}
}
