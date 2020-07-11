package main

import "fmt"

type data struct{
	name string
	age int
}

//reference: https://github.com/ardanlabs/gotraining/blob/master/topics/go/language/methods/example3/example3.go


func (d data) display(){
	fmt.Printf("Name is %s and age is %d",d.name,d.age)

}

func (d *data) increaseData(){
	d.age++
}

func main(){

 giri := data{
	 name : "Girimurugan", age : 39,
 }

giri.display()

giri.increaseData()
fmt.Println("====")

giri.display()


}