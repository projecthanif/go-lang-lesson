package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("File section")

	var content string = "My Name is Mustapha"

	file, err := os.Create("./myName.txt")

	if err != nil {
		myPanic(err)
	}

	len, err := io.WriteString(file, content)

	if err != nil {
		myPanic(err)
	}

	fmt.Println("Length is", len)
	defer file.Close()

	readFile("./myName.txt")
}

func myPanic(err error) {
	panic(err)
}

func readFile(filename string) {
	databyte, err := os.ReadFile(filename)

	if err != nil {
		panic(err)
	}

	fmt.Println("Text data inside the file is \n", databyte)
}
