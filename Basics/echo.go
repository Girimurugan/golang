package main

import (
	"fmt"
	"os"
)

func main() {

	var out, separator string

	for i := 1; i < len(os.Args); i++ {

		out += separator + os.Args[i]
		separator = " "

	}

	fmt.Println(out)

}
