package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Let play a game of number choose from 1 to 6")

	numGame()
}

func numGame() {
	bufioVal := bufio.NewReader(os.Stdin)
	input, _ := bufioVal.ReadString('\n')

	inputToInt, _ := strconv.ParseInt(strings.TrimSpace(input), 36, 64)

	if inputToInt < 1 || inputToInt > 6 {
		fmt.Println("Your chosen number cannot be less than 0 and greater than 6")
		numGame()
		return
	}

	switch inputToInt {
	case 1:
		fmt.Println("Open")
		
	case 2:
		fmt.Println("Congrats you won")
	case 3:
		fmt.Println("Try again")
	case 4:
		fmt.Println("Way to go")
	case 5:
		fmt.Println("Woah")
	case 6:
		fmt.Println("Unfortunately you thought that what you would hear")
	default:
		fmt.Println("What could have gone wrong")
	}

	if inputToInt == 2 {
		return
	}

	fmt.Println("Do you want to try again if so type `yes`")

	conti, _ := bufioVal.ReadString('\n')
	trimedAnswer := strings.TrimSpace(conti)

	if trimedAnswer == "yes" {
		numGame()
		return
	} else {
		fmt.Println("Goodbye!!!!")
	}

}
