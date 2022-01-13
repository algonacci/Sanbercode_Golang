package main

import "fmt"

var DatabaseConnection string

func init() {
	DatabaseConnection = "MySQL"
	fmt.Println("Melakukan init")
}
