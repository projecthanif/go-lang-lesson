package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type ApiResponse[T any] struct {
	Message string `json:"message"`
	Data    []T    `json:"data"`
}

type Service struct {
	Id          int          `json:"id"`
	Name        string       `json:"name"`
	Description string       `json:"description"`
	Subservices []Subservice `json:"subservices"`
}

type Subservice struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	ServiceId   int    `json:"serviceId"`
	ServiceName string `json:"serviceName"`
}

const GETBASEURL string = "https://my-fixer-api.laravel.cloud/api/services"

func main() {
	getRequestExample()
}

func getRequestExample() {
	response, err := http.Get(GETBASEURL)

	errorChecker(err)

	defer response.Body.Close()

	databytes, _ := io.ReadAll(response.Body)

	var apiRes ApiResponse[Service]

	json.Unmarshal(databytes, &apiRes)

	fmt.Println(apiRes.Message)
}

func errorChecker(err error) {
	if err != nil {
		panic(err)
	}
}
