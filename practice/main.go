package main

import (
	"fmt"
	"sort"
)

func main() {
	var colors = []string{"red", "green", "blue"}
	fmt.Println(colors)
	colors = append(colors, "purple")
	fmt.Println(colors)
	colors = append(colors[1:len(colors)])
	fmt.Println(colors)

	nubmers := make([]int, 5)
	nubmers[0] = 134
	nubmers[1] = 72
	nubmers[2] = 237
	nubmers[3] = 736
	nubmers[4] = 123
	fmt.Println(nubmers)

	sort.Ints(nubmers)
	fmt.Println(nubmers)

}
