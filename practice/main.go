package main

import (
	"fmt"
)

func main() {
	anInt := 42
	var p = &anInt
	fmt.Println(*p)

	value1 := 42.13
	pointer1 := &value1

	fmt.Println(*pointer1)
	*pointer1 = *pointer1 / 31
	fmt.Println(*pointer1)
	fmt.Println(value1)
}
