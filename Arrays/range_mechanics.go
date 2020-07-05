package main

import "fmt"

func main(){

	fruits := make([]string,3,3)


	fruits[0] = "Apple"
	fruits[1] = "Orange"
	fruits[2] = "Bums"

// the value semantic based ranges... in this the range copies the main slice and works on the copu

for _ , v:= range fruits{
	fruits = fruits[:1]
	fmt.Printf("v[%s]\n", v)
}

fmt.Printf("\n\n\n\n")

// iterations by the pointer semantics
for i := range fruits{
	fruits = fruits[:1]
	fmt.Printf("v[%s]\n", fruits[i])
}

}