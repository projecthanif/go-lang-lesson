package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

const url = "https://lco.dev"

func main() {

	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	databtyes, err := io.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	content := string(databtyes)

	contents := strings.Split(content, "go")

	fmt.Println(len(contents))

}
