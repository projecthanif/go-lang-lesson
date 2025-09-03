package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"syscall"

	"golang.org/x/term"
)

const RegUrl string = "https://my-fixer-api.laravel.cloud/api/auth/register"

type SuccessApiResponse struct {
	Message string `json:"message"`
	User    User   `json:"user"`
}

type User struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
	Role    string `json:"role"`
	Address string `json:"address"`
}

type Error struct {
	Message string              `json:"message"`
	Errors  map[string][]string `json:"errors"`
}

func main() {
	handleApiRequest()
}

func handleApiRequest() {
	data, err := handleUserInput()
	errorChecker(err)
	req, err := http.NewRequest("POST", RegUrl, strings.NewReader(data))
	errorChecker(err)

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	client := &http.Client{}

	res, err := client.Do(req)
	errorChecker(err)
	defer res.Body.Close()

	switch res.StatusCode {
	case 201:
		handleSuccessState(res)
	case 422:
		handleErrorState(res)
	default:
		handleOtherState(res)
	}

}

func handleSuccessState(body *http.Response) {
	var apiRes SuccessApiResponse
	data, err := io.ReadAll(body.Body)
	errorChecker(err)

	err = json.Unmarshal(data, &apiRes)
	errorChecker(err)

	fmt.Println("✅ Registration successful!")
	fmt.Println("Message:", apiRes.Message)
	fmt.Printf("User Details:\n")
	fmt.Printf("  - ID: %d\n", apiRes.User.Id)
	fmt.Printf("  - Name: %s\n", apiRes.User.Name)
	fmt.Printf("  - Email: %s\n", apiRes.User.Email)
	fmt.Printf("  - Phone: %s\n", apiRes.User.Phone)
	fmt.Printf("  - Role: %s\n", apiRes.User.Role)
	fmt.Printf("  - Address: %s\n", apiRes.User.Address)

}

func handleErrorState(body *http.Response) {
	var apiRes Error
	data, err := io.ReadAll(body.Body)
	errorChecker(err)

	err = json.Unmarshal(data, &apiRes)
	errorChecker(err)

	fmt.Println(apiRes.Message)

	errors := apiRes.Errors
	for i := range errors {
		ll(errors[i])
	}

	os.Exit(1)
}

func handleOtherState(body *http.Response) {
	ll("Status Code of" + strconv.Itoa(body.StatusCode))
	os.Exit(1)
}

func handleUserInput() (string, error) {

	inputs := make(map[string]string)

	inputs["role"] = "user"

	fields := []string{"name", "email", "password", "password_confirmation", "phone", "address"}

	for i := range fields {
		reader := bufio.NewReader(os.Stdin)

		var input string
		var err error

		fmt.Println("Input your", replaceUnderScore(capsWord(fields[i])))

		if fields[i] == "password" || fields[i] == "password_confirmation" {
			var bodyBtyes []byte
			bodyBtyes, err = term.ReadPassword(int(syscall.Stdin))

			input = strings.TrimSpace(string(bodyBtyes))
			inputs[fields[i]] = input
			fmt.Println("")
		} else {
			input, err = reader.ReadString('\n')
			input = strings.TrimSpace(string(input))
			inputs[fields[i]] = input
		}

		if err != nil {
			return "", err
		}
	}

	data, err := json.Marshal(inputs)

	return string(data), err
}

func errorChecker(err error) {
	if err != nil {
		panic(err)
	}
}

func capsWord(word string) string {
	first := string(word[0])

	return strings.ToUpper(first) + strings.SplitAfter(word, first)[1]
}

func replaceUnderScore(word string) string {
	return strings.ReplaceAll(word, "_", " ")
}

func ll(T any) {
	fmt.Println("ll ====>", T)
}
