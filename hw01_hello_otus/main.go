package main

import (
	"fmt"

	"golang.org/x/example/hello/reverse"
)

func main() {

	var stringToReverse = "Hello, OTUS!"

	var reversedString = reverse.String(stringToReverse)
	fmt.Println(reversedString)
}
