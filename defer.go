package main

import (
	"fmt"
)

func main() {
	defer fmt.Println("Close the file")
	fmt.Println("Open the file")

	defer fmt.Println("Statement 1")
	defer fmt.Println("Statement 2")

	myFunc()

	defer fmt.Println("Statement 3")
	fmt.Println("Statement 4")

	x := 1000
	defer fmt.Println("Valur of x:", x)
	x++
	fmt.Println("Value of x after incrementing: ", x)
}

func myFunc() {
	defer fmt.Println("Deferred in the function")
	fmt.Println("Not deffered in the function")
}