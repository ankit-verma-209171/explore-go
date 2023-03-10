package main

import "fmt"

func main() {
	fmt.Println("Slies in Go!")
	var fruits = []string{}

	fruits = append(fruits, "cherry")
	fmt.Println("Fruits:", fruits)
	fmt.Printf("Fruits type is %T \n", fruits)

	// another way to initialize slice
	nums := make([]int, 1)
	nums = append(nums, 1, 2, 3, 4, 5)
	fmt.Println("Nums are", nums)
	fmt.Println("Nums slice", nums[:3])

	// deleting element from slice
	colors := []string{"red", "blue", "yellow", "green"}
	fmt.Println("Colors", colors)

	fmt.Println("Removing yellow")
	index := 2

	colors = append(colors[:index], colors[index+1:]...)
	fmt.Println("Colors now", colors)
}
