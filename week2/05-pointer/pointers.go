package main

import "fmt"

type Student struct {
	Name    string
	IsAsdos bool
}

func passByValue(s Student) {
	s.Name = "New name"
}

func passByReference(s *Student) {
	s.IsAsdos = true
}

func main() {
	std := Student{
		Name:    "John Doe",
		IsAsdos: false,
	}

	passByValue(std)
	fmt.Println("student name:", std.Name)

	passByReference(&std)
	fmt.Println("is student an asdos:", std.IsAsdos)
}
