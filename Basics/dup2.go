package main

import (
	"fmt"
	"os"
	"bufio"

)

func main(){

	counts := make(map[string]int)

	filename := os.Args[1]

	fileReader, err := os.Open(filename)

	if err != nil {
		fmt.Println ("File Not Found")
	}

	input := bufio.NewScanner(fileReader)

	for input.Scan(){

		counts[input.Text()]++
	}

	for line, n := range counts{

		fmt.Println(line, "->", n)
	}



}