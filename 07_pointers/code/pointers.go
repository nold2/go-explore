package main

import "fmt"

// func main() {
// 	var name string
// 	var namePointer *string
//
// 	name = "Beyonce"
// 	namePointer = &name
// 	nameValue := *namePointer
//
// 	fmt.Println("Name:", name)
// 	fmt.Println("Name *:", namePointer)
// 	fmt.Println("Name again:", nameValue)
//
// 	some := "Beyonce"
// 	something := &some
//
// 	fmt.Println(some)
// 	fmt.Println(something)
// }

// // ******************************************************

// func main() {
// 	var name string = "Beyonce"
// 	var namePointer *string = &name
// 	var nameValue = *namePointer

// 	fmt.Println("Name:", name)
// 	fmt.Println("Name *:", namePointer)
// 	fmt.Println("Name Value:", nameValue)

// }

// // ******************************************************

// func changeName(n *string) {
// 	*n = strings.ToUpper(*n)
// }
//
// func main() {
// 	name := "Elvis"
// 	changeName(&name)
// 	fmt.Println(name)
// }

// // ******************************************************

// Coordinates represent a latitude and longitude
type Coordinates struct {
	X, Y float64
}

var c = Coordinates{X: 10, Y: 20}

func main() {
	coordinateMemoryAddress := c
	coordinateMemoryAddress.X = 200

	fmt.Println(coordinateMemoryAddress)
}
