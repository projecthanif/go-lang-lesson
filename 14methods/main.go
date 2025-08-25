package main

import "fmt"

func main() {
	user := User{"projecthanif", "projecthanif.dev", true}

	user.GetStatus()
}

type User struct {
	Name     string
	Email    string
	IsActive bool
}

func (u User) GetStatus() {
	fmt.Println("Is User Active:", u.IsActive)
}
