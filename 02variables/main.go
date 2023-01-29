package main

import "fmt"

// global walrus operator not allowed
// jwtToken := "sjd3893hd3ehdn38njd83" //not allowed

// Public variables start with caps
const LoginToken = "jshdu3db8392d89d"

func main() {

	// strings
	var username string = "ankit"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	// booleans
	isLoggedIn := true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	// floats
	var smallFloat float32 = 3.14159265359
	var bigFloat float64 = 3.14159265359
	fmt.Println(smallFloat)
	fmt.Println(bigFloat)

	// default values
	var password string
	var luckyNumber int
	var shouldShowStatus bool
	fmt.Println(password)
	fmt.Println(luckyNumber)
	fmt.Println(shouldShowStatus)

	// implicit variables
	var website = "go.dev"
	fmt.Println(website)
	
	// no var style
	noOfUsers := 12340
	fmt.Println(noOfUsers)

	fmt.Println(LoginToken)
	
}
