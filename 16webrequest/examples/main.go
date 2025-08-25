package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Service struct {
	Id          int          `json:"id"`
	Name        string       `json:"name"`
	Description string       `json:"description"`
	SubServices []SubService `json:"subservices"`
}

type SubService struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	ServiceId   int    `json:"serviceId"`
	ServiceName string `json:"serviceName"`
}

type APIResponse struct {
	Message string    `json:"message"`
	Data    []Service `json:"data"`
}

func main() {
	request, err := http.Get("https://my-fixer-api.laravel.cloud/api/services")

	if err != nil {
		panic(err)
	}

	defer request.Body.Close()

	if request.StatusCode != 200 {
		panic("Failed to fetch services")
	}

	content, err := io.ReadAll(request.Body)

	if err != nil {
		panic(err)
	}

	var responseBody APIResponse
	err = json.Unmarshal(content, &responseBody)

	if err != nil {
		panic(err)
	}

	fmt.Println(responseBody.Message)

}
