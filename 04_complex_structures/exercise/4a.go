package main

import "fmt"

func average(args ...float32) float32 {
	var sum float32
	for _, val := range args {
		sum += val
	}
	return sum / float32(len(args))
}

func main() {
	fmt.Println(average(10.0, 11.0, 12.0))
}
