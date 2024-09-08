package main

import "fmt"

type User struct {
	Id         int
	Username   string
	IsVerified bool
}

func (u *User) verify() {
	u.IsVerified = true
}

func main() {
	u := User{
		Id:         1,
		Username:   "user1234",
		IsVerified: false,
	}

	fmt.Println(u)

	u.verify()
	fmt.Println(u)
}
