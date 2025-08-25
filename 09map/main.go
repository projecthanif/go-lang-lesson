package main

import "fmt"

type User struct {
	Id       int
	Name     string
	Email    string
	Password string
}

func main() {
	user := make(map[string]string)

	user["username"] = "projecthanif"
	fmt.Println(user["username"])

	var number = make(map[int]string)

	number[1] = "one"

	fmt.Println(number[1])
	fmt.Printf("Type of number var %T\n", number)

	users := make(map[int]User)

	users[1] = User{1, "Mustapha Ibrahim", "iamustapha213@gmail.com", "password"}

	fmt.Println(users[1].Name)

}
