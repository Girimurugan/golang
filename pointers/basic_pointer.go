package main

import "fmt"

func main(){

	counter := 10
	fmt.Println(counter)

	increment(&counter)
	fmt.Println(counter)

}

// Pointer semantics
// This allows for learning about the pointer semantics
func increment(val *int){

	*val++

}