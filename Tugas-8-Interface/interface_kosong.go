package main

import (
	"fmt"
	"strings"
)

func main() {
	var secret interface{}

	secret = "ethan hunt"
	fmt.Println(secret)

	secret = []string{"apple", "manggo", "banana"}
	fmt.Println(secret)

	secret = 12.4
	fmt.Println(secret)

	secret = 2
	var number = secret.(int) * 10
	fmt.Println(secret, "multiplied by 10 is :", number)

	secret = []string{"apple", "manggo", "banana"}
	var fruits = strings.Join(secret.([]string), ", ")
	fmt.Println(fruits, "is my favorite fruits")
}
