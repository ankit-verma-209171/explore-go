package main

import "fmt"

func main() {
	fmt.Println("Experimenting with the pointers!")

	var ptr *int
	fmt.Println("Value of ptr is", ptr) // will be nil initially

	// fmt.Printf("Type of ptr is %T \n", *ptr) // will panic

	myInt := 7
	ptr = &myInt

	fmt.Println("Value at myInt address is", &myInt) 
	fmt.Println("Value at ptr is", ptr) 
	fmt.Println("Value at ptr reference is", *ptr) 

	// modifying value of myInt using ptr
	*ptr *= 6 // 7 x 6 = 42 (new value)
 	fmt.Println("New value of myInt is", myInt)

}
