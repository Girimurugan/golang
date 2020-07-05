package main

import "fmt"

func main(){

	friends := [3]string{"Giri","Gobi","Ram"}

	fmt.Println(friends)

	for i := range friends{
		friends[1] = "sita"

		if i == 1{
			fmt.Printf("Aft[%s]\n",friends[1])
			fmt.Println(friends)
		}
	}

	friends = [3]string{"Giri","Gobi","Ram"}

	for i, v := range &friends{

		friends[1] = "sita"
		if i == 1{
			fmt.Printf("Aft[%s]\t Address[%p] ",v, &v)
		}

	}

}