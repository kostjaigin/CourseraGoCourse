package main

import (
	"fmt"
	"strings"
)

func main() {

	animals := make(map[string]Animal)

	for {
		fmt.Printf("> ")
		// scan user input
		var command, name, arg string
		fmt.Scanln(&command, &name, &arg)
		switch strings.ToLower(command) {
		case "newanimal":
			storeAnimal(name, arg, animals)
		case "query":
			queryAnimals(name, arg, animals)
		default:
			fmt.Println("Invalid command")
		}
	}
}

func storeAnimal(name, kind string, collection map[string]Animal) {
	var animal Animal
	switch strings.ToLower(kind) {
	case "cow":
		animal = &Cow{Name: name}
	case "bird":
		animal = &Bird{Name: name}
	case "snake":
		animal = &Snake{Name: name}
	default:
		fmt.Println("Invalid animal type. Available types: [cow, bird, snake]")
		return
	}
	collection[name] = animal
	fmt.Println("Created it!")
}

func queryAnimals(name, action string, collection map[string]Animal) {
	animal, present := collection[name]
	if !present {
		fmt.Printf("No animal with name %s stored\n", name)
		return
	}
	switch strings.ToLower(action) {
	case "eat":
		animal.Eat()
	case "move":
		animal.Move()
	case "speak":
		animal.Speak()
	default:
		fmt.Println("Invalid animal command. Available commands: [eat, move, speak]")
	}
}

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct {
	Name string
}

type Bird struct {
	Name string
}

type Snake struct {
	Name string
}

func (*Cow) Eat() {
	fmt.Println("grass")
}

func (*Cow) Move() {
	fmt.Println("walk")
}

func (*Cow) Speak() {
	fmt.Println("moo")
}

func (*Bird) Eat() {
	fmt.Println("worms")
}

func (*Bird) Move() {
	fmt.Println("fly")
}

func (*Bird) Speak() {
	fmt.Println("peep")
}

func (*Snake) Eat() {
	fmt.Println("mice")
}

func (*Snake) Move() {
	fmt.Println("slither")
}

func (*Snake) Speak() {
	fmt.Println("hsss")
}
