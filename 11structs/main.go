package main

import "fmt"

func main() {
	fmt.Println("Structs in golang")
	// No inheritance in golang; No super or parent
	hitesh := User{"Daksh", "daksh@go.dev", true, 18}

	fmt.Println(hitesh)

	fmt.Printf("hitesh details are: %+v\n", hitesh)
	fmt.Printf("Name is %v and email is %v\n", hitesh.Name, hitesh.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
