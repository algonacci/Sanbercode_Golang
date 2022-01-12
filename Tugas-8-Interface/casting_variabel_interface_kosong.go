package main

import (
	"fmt"
)

type person struct {
	name string
	age  int
}

func main() {
	var secret interface{} = &person{name: "wick", age: 27}
	var name = secret.(*person).name
	var age = secret.(*person).age
	fmt.Println(name)
	fmt.Println(age)
}
