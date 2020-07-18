package main

import (
	"fmt"
)

//Define the interface not as object but a bhevaiour of the concreate type
type Speaker interface {
	speak()
	// Define what this behaviour can do like a Speaker can Speak
}


//Now which of the concreate types can have this bevaior

type dog struct{

}

type human struct{

}

func (d dog) speak(){
	fmt.Println("Bow Bow")
}

func (h human) speak(){
	fmt.Println("Blah Blah")
}

func whatTheySpeak(s Speaker){
	s.speak()
}

func main() {

	d := dog{}
	h := human{}

	whatTheySpeak(d)
	whatTheySpeak(h)

}
