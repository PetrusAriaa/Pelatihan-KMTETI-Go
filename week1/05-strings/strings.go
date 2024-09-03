package main

import "fmt"

func main() {
	// Working with string
	backendHost := "api.nightlogin.id"
	uri := "https://" + backendHost + "/users"
	fmt.Println(uri)

	uri2 := fmt.Sprint("https://", backendHost, "/users/1") // using Sprint
	fmt.Println(uri2)

	endpointName := "auth"
	uri3 := fmt.Sprintf("https://%s/%s", backendHost, endpointName) // using Sprintf
	fmt.Println(uri3)
}
