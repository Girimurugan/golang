package main

import ("fmt")

func main() {

	var fruits [5]string
	fruits[0] = "Apple"
	fruits[1] = "Orange"
	fruits[2] = "Papaya"
	fruits[3] = "Suppota"
	fruits[4] = "Lemon"

	for i, fruit := range fruits {
		fmt.Println(i, fruit)
	}

	pincodes := [5]int{1,2,3,4,5}

	for i:=0; i < len(pincodes); i++{

		fmt.Println(i,pincodes[i])
	}

}
