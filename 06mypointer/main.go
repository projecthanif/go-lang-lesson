package main

import "fmt"

func main() {
	num := 20
	num1 := 20
	num3 := 20
	var ptr1 *int = &num
	var ptr2 *int = &num1
	var ptr3 *int = &num3

	fmt.Printf("This address %v holds this value %v \n", &ptr1, *ptr1)
	fmt.Printf("This address %v holds this value %v \n", &ptr2, *ptr2)
	fmt.Printf("This address %v holds this value %v \n", &ptr3, *ptr3)

	fmt.Printf("This value %v, num 3 %v \n ", ptr3, &num3)

}
