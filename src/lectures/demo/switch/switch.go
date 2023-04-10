package main

import "fmt"

func price() int {
	return 1
}

const (
	Economy    = 0
	Business   = 1
	FirstClass = 2
)

func main() {
	switch p := price(); {
	case p < 2:
		fmt.Println("Cheap item")
	case p < 10:
		fmt.Println("Moderate item")
	default:
		fmt.Println("Expensive item")
	}
}
