package main

import "fmt"

func main() {
	array := []int{1, 2, 3, 4, 5}

	fmt.Printf("This array contains integer with length %v, capacity %v \n ", len(array), cap(array))
	fmt.Println(array)

}
