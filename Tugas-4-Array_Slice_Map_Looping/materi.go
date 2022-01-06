package main

import "fmt"

func main() {

	var names [4]string
	names[0] = "John"
	names[1] = "Doe"
	names[2] = "Frank"
	names[3] = "Jack"

	fmt.Println(names[0], names[1], names[2], names[3])

	var fruits = [4]string{
		"apple",
		"grape",
		"banana",
		"melon",
	}

	fmt.Println(fruits)

	var numbers = [...]int{2, 3, 2, 4, 3}

	fmt.Println("data array \t:", numbers)
	fmt.Println("jumlah elemen \t:", len(numbers))

	var fruits2 = []string{"apple", "grape", "banana", "melon"}
	fmt.Println(fruits2[0]) // "apple"
}
