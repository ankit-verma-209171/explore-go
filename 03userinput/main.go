package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the stars you want to give to Go😎:")

	// comma ok || error ok syntax
	input, _ := reader.ReadString('\n') //_ is used for values we don't care about
	fmt.Println("Thanks for the stars:", input)
}