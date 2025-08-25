package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://jsonplaceholder.typicode.com/comments?postId=1"

func main() {
	fmt.Println("Handling urls in golang")
	fmt.Println(myurl)

	result, _ := url.Parse(myurl)

	fmt.Println("Scheme: ", result.Scheme)
	fmt.Println("Host: ", result.Host)
	fmt.Println("Path: ", result.Path)
	fmt.Println("Port: ", result.Port())
	fmt.Println("Raw Query: ", result.RawQuery)

	qparams := result.Query()

	fmt.Printf("The Type of query params are: %T\n", qparams)
	fmt.Println(qparams["postId"])

	for _, val := range qparams {
		fmt.Println("Param is: ", val)
	}

}
