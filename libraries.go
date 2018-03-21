package main

import (
	"fmt"
	"stringutil"
)

func main() {

	n1, l1 := stringutil.Fullname("Kevin", "Peterson")
	fmt.Printf("Fullname: %v\n number of chars: %v\n", n1, l1)
	n2, l2 := stringutil.FullnameNakedReturn("Hardik", "Pandya")
	fmt.Printf("Fullname: %v\n number of chars: %v\n", n2, l2)
}