package main

import (
	"fmt"
)

func main() {
	var colors [3]string
	colors[0] = "Red"
	colors[1] = "Green"
	colors[2] = "Blue"
	fmt.Println(colors)
	fmt.Println(colors[2])

	var numbers = [5]int{1, 2, 3, 4, 5}
	fmt.Println(numbers)

	fmt.Println("Length of colors array is", len(colors))
}
