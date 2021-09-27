package main

import (
	math "fem-intro-to-go/05_toolkit/code/utils"
	"fmt"
)

func calculate() int {
	total := math.Add(1, 2, 3)
	return total
}

func main() {
	fmt.Println(calculate())
	fmt.Println("Packages!")
}
