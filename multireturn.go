package main

import (
	"fmt"
)

func main() {

	n1, l1 := Fullname("Kevin", "Peterson")
	fmt.Printf("Fullname: %v\n number of chars: %v\n", n1, l1)
	n2, l2 := FullnameNakedReturn("Hardik", "Pandya")
	fmt.Printf("Fullname: %v\n number of chars: %v\n", n2, l2)
}

func Fullname(f, l string)(string, int) {
	full := f + " " + l
	length := len(full)
	return full, length
}

func FullnameNakedReturn(f, l string)(full string,length int) {
	full = f + " " + l
	length = len(full)
	return
}