package main

import "fmt"

type User struct {
	Id    uint8
	Name  string
	Email string
}

func main() {
	// declare as a reference

	u0 := &User{
		Id: 0, Name: "Michael", Email: "michael123@example.com",
	}
	fmt.Println(u0)

	u1 := User{
		Id: 0, Name: "John", Email: "john@example.com",
	}
	u2 := User{
		Id: 1, Name: "Emma", Email: "emma@example.com",
	}
	users := []*User{&u1, &u2}

	fmt.Println(users)
	fmt.Println(users[0])
}
