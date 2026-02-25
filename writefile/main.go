package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please supply filename to display as a first argument")
		os.Exit(1)
	}
	fileName := os.Args[1]

	f, e := os.Open(fileName)

	if e != nil {
		fmt.Println("Error occured", e)
		os.Exit(1)
	}
	io.Copy(os.Stdout, f)
}
