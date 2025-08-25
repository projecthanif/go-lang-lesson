package main

import "fmt"

func main() {
	defer greeter()
	fmt.Println("Hello from main")
}

func greeter() {
	fmt.Println("Hello from greeter")
}
