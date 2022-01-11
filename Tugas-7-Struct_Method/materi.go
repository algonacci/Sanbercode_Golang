package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	var john = struct {
		person
		grade int
	}{}
	john.person = person{"wick", 21}
	john.grade = 2

	fmt.Println("name  :", john.person.name)
	fmt.Println("age   :", john.person.age)
	fmt.Println("grade :", john.grade)
}
