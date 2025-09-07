package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	// performGetRequest()
	performPostRequest()
}

func performPostRequest() {
	fmt.Println("This is a post request")
	const myUrl string = "https://my-fixer-api.laravel.cloud/api/auth/login"

	requestBody := strings.NewReader(`
		{
			"email":"artisantest1@gmail.com",
			"password":"password
		}
	`)

	res, err := http.Post(myUrl, "application/json", requestBody)

	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	content, _ := io.ReadAll(res.Body)
	fmt.Println(string(content))
}

func performGetRequest() {
	fmt.Println("This is a get request")
	const myUrl string = "https://jsonfakery.com/users/random"

	res, err := http.Get(myUrl)

	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		panic("Wtf!!!!")
	}

	var responseString strings.Builder
	byteData, _ := io.ReadAll(res.Body)
	btyeCount, _ := responseString.Write(byteData)

	fmt.Println(btyeCount)

	fmt.Println(responseString.String())

}
