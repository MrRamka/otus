package main

import (
	"fmt"

	"golang.org/x/example/hello/reverse"
)

func main() {
	stringToReverse := "Hello, OTUS!"
	reversedString := reverse.String(stringToReverse)
	fmt.Println(reversedString)
}
