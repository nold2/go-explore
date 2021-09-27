package main

import "fmt"

// User is a user type
type User struct {
	ID                         int
	FirstName, LastName, Email string
}

// // User is a user type
// type User struct {
// ID					       int
// FirstName, LastName, Email string
// }

//Group represent a set of users
type Group struct {
	Role           string
	Users          []User
	lastUser       User
	spaceAvailable bool
}

func (u *User) describe() string {
	desc := fmt.Sprintf("Name: %s %s, Email: %s", u.FirstName, u.LastName, u.Email)
	return desc
}

func describe(u User) string {
	desc := fmt.Sprintf("Name: %s %s, Email: %s", u.FirstName, u.LastName, u.Email)
	return desc
}

func (g *Group) describe() string {
	g.spaceAvailable = !(len(g.Users) > 2)
	desc := fmt.Sprintf("This user group has %d users. The newest user is %s %s. Accepting New Users: %t", len(g.Users), g.lastUser.FirstName, g.lastUser.LastName, g.spaceAvailable)

	return desc
}

//func describeGroup
// => "This user group has 18 users. The newest user is Joe Smith. Accepting New Users: true"
func describeGroup(g Group) string {
	g.spaceAvailable = !(len(g.Users) > 2)
	desc := fmt.Sprintf("This user group has %d users. The newest user is %s %s. Accepting New Users: %t", len(g.Users), g.lastUser.FirstName, g.lastUser.LastName, g.spaceAvailable)

	return desc
}

func main() {
	u := User{ID: 1, FirstName: "Marilyn", LastName: "Monroe", Email: "marilyn.monroe@gmail.com"}
	u2 := User{ID: 2, FirstName: "John", LastName: "Doe", Email: "john.doe@gmail.com"}
	u3 := User{ID: 2, FirstName: "John", LastName: "Doe", Email: "john.doe@gmail.com"}

	g := Group{
		Role:           "admin",
		Users:          []User{u, u2, u3},
		lastUser:       u2,
		spaceAvailable: true,
	}

	fmt.Println(describe(u))
	fmt.Println(describeGroup(g))

	fmt.Println(u.describe(), g.describe())
}
