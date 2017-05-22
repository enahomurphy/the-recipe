package main

import (
	"fmt"
	"os"
)

func main() {
	// cannot be reassigned
	const hello = "Hello"

	// declare a variable in golan
	var world string
	//others
	// int
	// int32
	// int64 this three integers also have a float counterpart

	// assigning a variable
	world = "guest"

	// go offers dynamic assignment
	geetings := " have a wonder full weekend"
	if len(os.Args) > 1 {
		fmt.Println(hello + " " + os.Args[1] + geetings)
	} else {
		fmt.Println(hello + " " + world + geetings)
	}
}
