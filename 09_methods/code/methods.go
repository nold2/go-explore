package main

import "fmt"

// User is a user type
type User struct {
	ID                         int
	FirstName, LastName, Email string
}

func (u *User) describe() string {
	return fmt.Sprintf("Name: %s %s, Email: %s", u.FirstName, u.LastName, u.Email)
}

func describeUser(u User) string {
	desc := fmt.Sprintf("Name: %s %s, Email: %s", u.FirstName, u.LastName, u.Email)
	return desc
}

func main() {
	user := User{ID: 1, FirstName: "Marilyn", LastName: "Monroe", Email: "marilyn.monroe@gmail.com"}

	desc := describeUser(user)
	fmt.Println(desc)

	fmt.Println(user.describe())
}
