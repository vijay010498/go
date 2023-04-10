package main

import "fmt"

func main() {
	var myName = "vijay"
	fmt.Println("my name is", myName, myName)

	var name string = "pablo"
	fmt.Println("name = ", name)

	userName := "admin"
	fmt.Println("userName = ", userName)

	var sum int
	fmt.Println("sum = ", sum)

	part1, other := 1, 5
	fmt.Println("part 1 is ", part1, "other is ", other)

	part2, other := 2, 0
	fmt.Println("part 2 is ", part2, "other is ", other)

	sum = part1 + part2
	fmt.Println("sum = ", sum)

	var (
		lessonName = "variables"
		lessonType = "Demo"
	)
	fmt.Println("lessonName = ", lessonName, "type = ", lessonType)

	word1, word2, _ := "hello", "world", "!"

	fmt.Println("word1 = ", word1, "word2 = ", word2);
}
