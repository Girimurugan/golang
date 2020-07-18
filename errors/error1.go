package main

import (
	"errors"
	"fmt"
)

var (
	ErrBackendError = errors.New("Backend Error")
	ErrPageMoved    = errors.New("Page Moved")
)

func main() {

	if err := webCall(); err != nil {
		switch err {
		case ErrBackendError:
			fmt.Println("Backend Error")
		case ErrPageMoved:
			fmt.Println("Page Moved Error")
		}
	}

}

func webCall() error {
	return ErrBackendError
}
