package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func main() {
	performFormdataRequest()
}

func performFormdataRequest() {

	const myUrl = "https://my-fixer-api.laravel.cloud/api/auth/login"
	data := url.Values{}

	data.Add("email", "iamustapha213@gmail.com")
	data.Add("password", "password")

	res, err := http.PostForm(myUrl, data)

	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	content, _ := io.ReadAll(res.Body)

	fmt.Println(string(content))
}
