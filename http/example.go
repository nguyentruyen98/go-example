package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("asdw")
	f, err := os.Open(os.Args[1]) // For read access.
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	io.Copy(os.Stdout, f)

}
