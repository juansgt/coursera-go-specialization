package main

import (
	"fmt"
)

type Animal struct {
	food       string
	locomotion string
	noise      string
}

func NewAnimal(food, locomotion, noise string) *Animal {
	animal := &Animal{
		food:       food,
		locomotion: locomotion,
		noise:      noise,
	}

	return animal
}

func (animal *Animal) Eat() string {
	return animal.food
}

func (animal *Animal) Move() string {
	return animal.locomotion
}

func (animal *Animal) Speak() string {
	return animal.noise
}

func main() {
	var requestedAnimal string
	var requestedInformation string
	var animal *Animal

	requestedAnimal = "0"
	fmt.Println("Please insert following command: <animal> <information> Ex: cow move")

questionLoop:
	for {
		fmt.Print("> ")
		fmt.Scan(&requestedAnimal)

		switch requestedAnimal {
		case "cow":
			animal = NewAnimal("grass", "walk", "moo")
		case "bird":
			animal = NewAnimal("worms", "fly", "peep")
		case "snake":
			animal = NewAnimal("mice", "slither", "hsss")
		case "X":
			break questionLoop
		default:
			fmt.Println("Don't know that animal")
		}

		if animal != nil {

			fmt.Scan(&requestedInformation)

			switch requestedInformation {
			case "eat":
				fmt.Println(animal.Eat())
			case "move":
				fmt.Println(animal.Move())
			case "speak":
				fmt.Println(animal.Speak())
			default:
				fmt.Println("Don't know that information about that animal")
			}
		}
	}
}
