package main

import (
	"fmt"
	"time"
)

//time package - after, sleep, timelocation, day, date, ns, month
//timeformat - used regularly, provides syntax for formating

func main() {
	fmt.Println("Welcome to Time Study of Golang")

	presentTime := time.Now()
	fmt.Println(presentTime)

	//different format
	//Always use this date, time, and Monday as standard format
	fmt.Println(presentTime.Format("01-02-2006"))
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))

	//only month is of type time.Month
	createDate := time.Date(2020, time.August, 22, 06, 0, 0, 0, time.UTC)
	fmt.Println(createDate)
	fmt.Println(createDate.Format("01-02-2006 Monday"))

	// go env
	// GOOS = "darwin"
	//build for different OS
	//GOOS="windows" go build
}
