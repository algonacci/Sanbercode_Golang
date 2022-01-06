package main

import "fmt"

func main() {
	// Soal 1
	fmt.Println("------ SOAL KE 1 ------")
	for i := 1; i <= 20; i++ {
		if i%2 == 0 {
			fmt.Println("Berkualitas")
		} else if i%3 == 0 {
			fmt.Println("I Love Coding")
		} else {
			fmt.Println("Santai")
		}
	}
	// Soal 2
	fmt.Println("------ SOAL KE 2 ------")
	for x := 0; x <= 7; x++ {
		for y := 0; y < x; y++ {
			fmt.Printf("#")
		}
		fmt.Println()
	}

	// Soal 3
	fmt.Println("------ SOAL KE 3 ------")
	var kalimat = [...]string{"aku", "dan", "saya", "sangat", "senang", "belajar", "golang"}
	fmt.Println(kalimat[2:])

	// Soal 4
	fmt.Println("------ SOAL KE 4 ------")
	var sayuran = []string{
		"Bayam",
		"Buncis",
		"Kangkung",
		"Kubis",
		"Seledri",
		"Tauge",
		"Timun",
	}
	for i := 0; i < 7; i++ {
		fmt.Println(i+1, sayuran[i])
	}

	// Soal 5
	fmt.Println("------ SOAL KE 5 ------")
	var satuan = map[string]int{
		"panjang": 7,
		"lebar":   4,
		"tinggi":  6,
	}
	var volumebalok = 1
	for i, data := range satuan {
		fmt.Println(i, "=", data)
		volumebalok = data * volumebalok
	}
	fmt.Println("Volume Balok", "=", volumebalok)

}
