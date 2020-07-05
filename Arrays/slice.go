package main

import "fmt"

func main(){

	fruits := make([]string,3,5)

	fruits[0] = "Apple"
	fruits[1] = "Orange"


	for i := range fruits{
		fmt.Println(fruits)
		fmt.Println(i)
	}

	fmt.Printf("Length [%d] and Capacity is [%d]", len(fruits), cap(fruits))
	fmt.Println(&fruits[2])
}