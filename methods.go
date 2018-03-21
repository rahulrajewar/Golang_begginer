package main

import (
	"fmt"
)

type Dogg struct {
	Breed string
	Weight int
	Sound string
}

func (d Dogg) speak() {
	fmt.Println(d.Sound)
}

func (d *Dogg) SpeakThreeTimes() {
	d.Sound = fmt.Sprintf("%v! %v! %v!", d.Sound,d.Sound, d.Sound)
	fmt.Println(d.Sound)
}

func main() {
	poodle := Dogg{"Poodle", 37, "Woof"}
	fmt.Println(poodle)
	poodle.speak()

	poodle.Sound = "Arf"
	poodle.speak()

	poodle.SpeakThreeTimes()
	poodle.SpeakThreeTimes()
}