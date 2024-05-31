package main

import (
	"fmt"
)

const aConst int = 64

func main() {
	var aString string = "The format is so strange!"
	fmt.Println(aString)
	fmt.Printf("The variable is of type: %T\n", aString)

	var anInteger int = 42
	fmt.Println(anInteger)

	var anotherString = "This is anothe string"
	fmt.Println(anotherString)
	fmt.Printf("The variable is of type: %T\n", anotherString)

	myString := "This is a string too"
	fmt.Println(myString)
	fmt.Printf("The variable is of type: %T\n", myString)
}
