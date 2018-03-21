package main

import (
	"fmt"
)

func main() {
	dosomething()

	sum := addvalues(23,54)
	fmt.Println("sum:", sum)

	sum = addAllvalues(12,32,56,10,20,20)
	fmt.Println("New sum:", sum)
}

func dosomething() {
	fmt.Println("Doing something!")
}

func addvalues(value1, value2 int)int {
	return value1+value2
}

func addAllvalues(values ...int) int {
	sum := 0
	for i:= range values{
		sum += values[i]
	}
	fmt.Printf("%T\n", values)
	return sum
}