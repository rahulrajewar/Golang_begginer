package main

import (
	"fmt"
)

func main() {

	str1 := "The quick red fox"
	str2 := "Jumped over"
	str3 := "the lazy brown dog"
	aNumber := 42
	isTrue := true
	stringLength, err := fmt.Println(str1,str2,str3)

	if err == nil {
		fmt.Println("String Length: ", stringLength)
	}

	fmt.Printf("Value of aNumber = %v\n", aNumber)
	fmt.Printf("Value of isTrue = %v\n", isTrue)
	fmt.Printf("Valur of aNumber as a Float = %.2f\n", float64(aNumber))

	fmt.Printf("Data Types: %T, %T, %T, %T, and %T\n", str1, str2, str3, aNumber, isTrue)
	myString := fmt.Sprintf("Data Types as var: %T, %T, %T, %T, and %T\n", str1, str2, str3, aNumber, isTrue)
	fmt.Print(myString)
}
