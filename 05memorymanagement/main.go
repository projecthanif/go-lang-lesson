package main

import (
	"fmt"
	"runtime"
)

func main() {
	memory := runtime.NumCPU()
	os := runtime.GOOS

	fmt.Printf("This PC has %v number of cpu \n", memory)
	fmt.Printf("This PC is running %v OS\n", os)
}
