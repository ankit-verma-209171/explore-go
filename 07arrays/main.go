package main

import "fmt"

func main() {
	fmt.Println("Arrays in Go")

	var fruitList [4]string
	fruitList[0] = "Apple"
	fruitList[2] = "Banana"

	fmt.Println("fruit list ->", fruitList)
	fmt.Println("fruit list size ->", len(fruitList))

	vegList := [5]string{"potato"}
	fmt.Println("vegetable at index 0:", vegList[0])
	fmt.Println("vegetable at index 0:", vegList[1])
}