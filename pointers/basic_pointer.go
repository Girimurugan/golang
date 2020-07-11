package main

import "fmt"

func main(){

	counter := 10
	fmt.Println(counter)

	increment(&counter)
	fmt.Println(counter)

}

func increment(val *int){

	*val++

}