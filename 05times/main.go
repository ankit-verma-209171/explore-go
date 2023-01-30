package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Experimenting with time")
	now := time.Now()

	fmt.Println("Now the time is", now)
	formattedTime := now.Format("2006/01/02 Monday 15:04:02")
	fmt.Println("Formatted Time:", formattedTime)

	twoDaysAfter := now.Add(time.Hour * 24 * 2)
	fmt.Println("2 days after will be", twoDaysAfter.Format("Monday"))
}
