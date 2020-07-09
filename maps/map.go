package main

import "fmt"

type country struct{
	name string
	continent string
}

func main(){

 countries := make(map[string]country)

 // adding the country using the normal mapping
 countries["IN"] = country{"Indian","Asia"}
 countries["US"] = country{"USA", "America"}

 fmt.Println(countries)

 fmt.Println("-========")

 //adding countries using the string liteal format
 countries_l := map[string]country{

	 "IN": {"India","Asia"},
	 "US": {"USA","America"},
 }

 fmt.Println(countries_l)

 fmt.Println("-========")

 for key, value := range countries{
	 fmt.Println(key, value)

 }

 fmt.Println("-========")

 // the ok or found flag tells whether the map entry was found and it is not empty value
  country_in, ok := countries["IN"]
  fmt.Println(country_in,ok)


}