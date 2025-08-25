package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time study")

	presentTime := time.Now()

	fmt.Printf("The time is %v \n", presentTime)

	formatedDay := presentTime.Format("Monday")
	formatedDate := presentTime.Format(time.UnixDate)
	formatedMonth := presentTime.Format("January")
	var formatedYear time.Month = presentTime.Month()

	fmt.Println(time.UnixDate)
	fmt.Printf("Today's Date is %v \n", formatedDate)
	fmt.Printf("Today is %v \n", formatedDay)
	fmt.Printf("We are in %v \n", formatedMonth)
	fmt.Printf("We are in the year %v \n", formatedYear)

}
