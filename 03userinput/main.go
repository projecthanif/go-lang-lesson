package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Give us your rating")
	input, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("Error due to:", err)
		return
	}

	inputParseToInt, _ := strconv.ParseInt(strings.TrimSpace(input), 10, 64)

	if inputParseToInt > 5 || inputParseToInt < 0 {
		fmt.Println("Input can not be less than 0 and greater than 5")
		main()
		return
	}

	fmt.Println("Thank you for your rating:", inputParseToInt)

}
