package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"golang.org/x/term"
)

type ApiResponse struct {
	Message string `json:"message"`
	User    User   `json:"user"`
	Token   string `json:"token"`
}

type User struct {
	Id               int           `json:"id"`
	Name             string        `json:"name"`
	Email            string        `json:"email"`
	Phone            string        `json:"phone"`
	Role             string        `json:"role"`
	Address          string        `json:"address"`
	GeoCoordinate    GeoCoordinate `json:"geo_coordinates"`
	Avater           string        `json:"avater"`
	IsVerified       bool          `json:"is_verified"`
	ProfileCompleted bool          `json:"profileCompleted"`
}

type GeoCoordinate struct {
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
}

const LoginUrl string = "https://my-fixer-api.laravel.cloud/api/auth/login"

func main() {
	data := userInput()
	sendLoginRequest(data)
}

func sendLoginRequest(data string) {
	req, err := http.NewRequest("POST", LoginUrl, strings.NewReader(data))

	errorChecker(err)

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	client := &http.Client{}

	res, err := client.Do(req)
	errorChecker(err)
	defer res.Body.Close()

	switch res.StatusCode {
	case 200:
		handleSuccess(res)
	default:
	}
}

func handleSuccess(body *http.Response) {
	var apiRes ApiResponse
	res, err := io.ReadAll(body.Body)
	errorChecker(err)

	json.Unmarshal(res, &apiRes)

	fmt.Println(apiRes.Message)

	fmt.Println(apiRes.User)
}

func userInput() string {
	reader := bufio.NewReader(os.Stdin)

	fields := []string{"email", "password"}

	inputs := make(map[string]string)

	fmt.Println("Input your credentails")

	for i := range fields {
		fmt.Println("Enter your", fields[i])
		if fields[i] == "password" {
			password, err := term.ReadPassword(int(os.Stdin.Fd()))
			errorChecker(err)
			inputs[fields[i]] = strings.TrimSpace(string(password))
		} else {
			dataBtye, err := reader.ReadString('\n')
			errorChecker(err)
			inputs[fields[i]] = strings.TrimSpace(string(dataBtye))
		}
	}

	jsonEncode, err := json.Marshal(inputs)
	errorChecker(err)
	return string(jsonEncode)
}

func errorChecker(err error) {
	if err != nil {
		fmt.Print("Error due to:")
		panic(err)
	}
}
