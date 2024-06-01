package main

import (
	"fmt"
	"time"
)

func main() {
	n := time.Now()
	fmt.Println("I recorded the time as", n)

	t := time.Date(2019, time.November, 10, 23, 0, 0, 0, time.UTC)
	fmt.Println("Go lunched at ", t)
	fmt.Println(t.Format(time.Kitchen))
}
