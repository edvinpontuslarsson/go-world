package main

import (
	"fmt"
)

func getMessage() string {
	return "Go World"
}

// go run main.go

func main() {
	var intro = "Message:"
	msg := getMessage()

	fmt.Println(intro, msg)
}
