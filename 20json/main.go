package main

import (
	"encoding/json"
	"fmt"
)

type job struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Company     string `json:"company"`
	Location    string `json:"location,omitempty"`
}

func main() {
	fmt.Println("Welcome to json video")
	jsonEncoded, _ := encodeJson()

	fmt.Println(string(jsonEncoded))

	decodeJson()
}

func encodeJson() ([]byte, error) {
	myJobs := []job{
		{"bb0a49cb-c719-4057-ab64-a77be8747457", "Personal Financial Advisor", "Personal Financial Advisor", "Braun Ltd", "Bacolod, Kiribati"},
		{"33933cf1-7a65-4158-a3fb-63788f3dd632", "Fast Food Cook", "McDermott, Jacobson and Hodkiewicz", "Braun Ltd", "Bacolod, Kiribati"},
	}

	return json.Marshal(myJobs)

}

func decodeJson() {
	jsonData := []byte(`{
    "id": "bb0a49cb-c719-4057-ab64-a77be8747457",
    "title": "Personal Financial Advisor",
    "description": "Minima repellat consequatur eum eius magni facilis. Quaerat maiores cum excepturi laborum est.",
    "company": "Braun Ltd",
    "location": "Bacolod, Kiribati"
  }`)

	isValid := json.Valid(jsonData)

	if !isValid {
		panic(isValid)
	}

	var jobRes job

	err := json.Unmarshal(jsonData, &jobRes)

	if err != nil {
		panic(err)
	}

	fmt.Println(jobRes)

}
