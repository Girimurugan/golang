package main

import "fmt"

func main(){

	counter := 10
	fmt.Println(counter)

	increment(&counter)
	fmt.Println(counter)

}

// Pointer semantics
func increment(val *int){

	*val++

}