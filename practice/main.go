package main

import (
	"fmt"
)

func main() {
	i1, i2, i3 := 12, 45, 68
	intSum := i1 + i2 + i3
	fmt.Print("Integer Sum: ", intSum, "\n")

	f1, f2, f3 := 23.5, 65.1, 76.3
	floatSum := f1 + f2 + f3
	fmt.Printf("Float Sum: %.2f\n", floatSum)

}
