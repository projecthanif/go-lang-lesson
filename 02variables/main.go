package main

import "fmt"

const LoginToken string = "Loroskak"

func main() {
	var username string = "projecthanif"
	fmt.Println(username)

	var isNewbie bool = true
	fmt.Println(isNewbie)

	var withOutDataType = "withOutDataType"
	fmt.Println(withOutDataType)

	withOutVarAndDataType := "withOutVarAndDataType"
	fmt.Println(withOutVarAndDataType)

	var anotherWay = "another way to create a variable with just `var`"
	fmt.Println(anotherWay)

	anotherWay2 := "another way to create a variable with `:`"
	fmt.Println(anotherWay2)

	fmt.Println(LoginToken)
}
