package main

import "fmt"

func printAge(age1, age2 int) (ageOfSally, ageOfBob int) {
	ageOfSally = age1
	ageOfBob = age2
	return
}

func printSome(ages ...int) int {
	var result int
	for _, value := range ages {
		result += value
	}
	return result
}

func main() {
	age1, age2 := printAge(20, 21)
	fmt.Println(age1, age2)
	fmt.Println(printSome(1, 2, 3))
}
