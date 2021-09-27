package main

import "fmt"

type User struct {
	ID                         int
	FirstName, LastName, Email string
}

func updateEmail(u *User, e string) {
	u.Email = e
}

func main() {
	user := User{ID: 1, FirstName: "ARNOLD", LastName: "Sebastian", Email: "arnold@sebastian.com"}
	updateEmail(&user, "hi@arnoldsebastian.com")
	fmt.Println(user)
}
