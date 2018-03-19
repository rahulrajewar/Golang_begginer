package main

import(
	"fmt"
	"sort"
)

func main() {

	var colors = []string{"Red", "Blue", "Green"}
	fmt.Println(colors)

	colors = append(colors, "Purple")
	fmt.Println(colors)

	colors = append(colors[1:len(colors)])
	fmt.Println(colors)

	colors = append(colors[:len(colors)-1])
	fmt.Println(colors)

	numbers := make([] int, 5, 5)
	numbers[0] = 152
	numbers[1] = 2
	numbers[2] = 32
	numbers[3] = 41
	numbers[4] = 53
	fmt.Println(numbers)

	numbers = append(numbers, 35)
	fmt.Println(numbers)
	fmt.Println(cap(numbers))

	sort.Ints(numbers)
	fmt.Println(numbers)
}