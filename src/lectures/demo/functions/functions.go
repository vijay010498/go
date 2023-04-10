package main

import "fmt"

func double(x int) int {
	return x + x
}

func add(lhs, rhs int) int {
	return lhs + rhs
}

func greet() {
	fmt.Println("Hello from go")
}

func main() {
	greet()

	dozen := double(6)
	fmt.Println("A dozen is", dozen)

	sum := add(2, 3)
	fmt.Println("Sum is", sum)

}
