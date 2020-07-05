package main

import "fmt"

func main() {

	// Declare an array of 5 integers that is initialized
	// to its zero value.
	var five [5]int

	// Declare an array of 4 integers that is initialized
	// with some values.
	four := [5]int{10, 20, 30, 40,50}

	// Assign one array to the other
	five = four

	// ./example2.go:21: cannot use four (type [4]int) as type [5]int in assignment

	fmt.Println(four)
	fmt.Println(five)
}