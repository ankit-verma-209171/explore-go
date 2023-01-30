package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	fmt.Println("Welcome to our pizza app")
	fmt.Printf("Please rate our pizza between 1 and 5: ")

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	fmt.Println("Thanks for the rating:", input)

	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		fmt.Println("Oops! Something went wrong :( \n", err)
	} else {
		numRating += 1
		fmt.Println("Num Ratings:", numRating)
	}
}
