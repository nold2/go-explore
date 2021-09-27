package main

import "fmt"

func main() {
	sentence := "I'm doing an exercise from frontendmasters"

	for index, val := range sentence {
		if index%2 == 1 {
			fmt.Println(string(val))
		}
	}
}
