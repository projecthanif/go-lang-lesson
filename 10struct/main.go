package main

import "fmt"

type User struct {
	Name       string
	Email      string
	IsVerified bool
	Age        int
	Role       UserRole
}

type UserRole struct {
	Name string
}

func main() {

	hanif := User{"Projecthanif", "projecthanif@hanif.me", true, 22, UserRole{"Admin"}}

	fmt.Println(hanif)
	// %+v  to give full context
	fmt.Printf("hanif info is as follows: %+v\n", hanif)

	midnightInn := User{}
	midnightInn.Name = "The Midnight Inn"
	midnightInn.Email = "support@midnightinn.io"
	midnightInn.IsVerified = true
	midnightInn.Age = 1
	midnightInn.Role.Name = "Admin"

	fmt.Printf("midnightInn info: %+v\n", midnightInn)

	if hanif.Role.Name != "Admin" {
		fmt.Println("You are UnAuthorized 401")
	} else {
		fmt.Println("You are Authorized 200")
	}

}
