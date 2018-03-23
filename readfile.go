package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	filename := "./fromString.txt"

	content, err := ioutil.ReadFile(filename)
	checkError(err)

	result := string(content)
	fmt.Println("Read from file:", content)
	fmt.Println("Read from file:", result)


}


func checkError(err error){
	if err != nil {
		panic(err)
	}
}