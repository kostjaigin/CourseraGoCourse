package main

import (
	"fmt"
	"strings"
)

func main() {

	animals := map[string]Animal{
		"cow":   CreateAnimal("grass", "walk", "moo"),
		"bird":  CreateAnimal("worms", "fly", "peep"),
		"snake": CreateAnimal("mice", "slither", "hsss"),
	}

	for {
		fmt.Printf(">")
		// scan user input
		var animalName, animalInfo string
		fmt.Scan(&animalName, &animalInfo)
		animal, animalPresent := animals[animalName]
		if animalPresent {
			switch strings.ToLower(animalInfo) {
			case "eat":
				fmt.Println(animal.Eat())
			case "move":
				fmt.Println(animal.Move())
			case "speak":
				fmt.Println(animal.Speak())
			default:
				fmt.Println("Invalid action")
			}
		} else {
			fmt.Println("Invalid animal")
		}
	}
}

type Animal struct {
	food, locomotion, noise string
}

func (a *Animal) init(food, locomotion, noise string) {
	a.food = food
	a.locomotion = locomotion
	a.noise = noise
}

func CreateAnimal(food, locomotion, noise string) Animal {
	var a Animal
	a.init(food, locomotion, noise)
	return a
}

func (a *Animal) Eat() string {
	return a.food
}

func (a *Animal) Move() string {
	return a.locomotion
}

func (a *Animal) Speak() string {
	return a.noise
}
