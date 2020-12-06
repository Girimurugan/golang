package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	out, separator := "", " "

	for i, arg := range os.Args[1:] {
		fmt.Println(i)
		out += strconv.Itoa(i) + separator + arg
	}

	fmt.Println(out)
	fmt.Println(os.Args[1:])
	fmt.Println(os.Args[0])
}
