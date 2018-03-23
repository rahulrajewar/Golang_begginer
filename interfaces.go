package main

import (
	"fmt"
)

type Animal interface {
	Speak( ) string
}

type Dogs struct {
}
func (d Dogs) Speak() string {
	return "Woof!"
}

type Cat struct {
}
func (d Cat) Speak() string {
	return "Meow!"
}

type Cow struct {
}
func (d Cow) Speak() string {
	return "Mooo!"
}

func main() {
	poodle:= Animal(Dogs{})
	fmt.Println(poodle)

	animals := []Animal{Dogs{}, Cat{}, Cow{}}
	for _, animal := range animals{
		fmt.Println(animal.Speak())
	}
}