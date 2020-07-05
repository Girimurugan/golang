package main

import "fmt"

type user struct{
	likes int
}

func main(){

	users := make([]user,3,3)

	sharedUser := &users[1]
	sharedUser.likes++

	fmt.Println(users[1].likes)
}