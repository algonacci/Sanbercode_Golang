package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// Soal 1
	fmt.Println("=========== soal no 1 ==============")
	nama1 := "Bootcamp"
	nama2 := "Digital"
	nama3 := "Skill"
	nama4 := "Sanbercode"
	nama5 := "Golang"

	fmt.Println(nama1, nama2, nama3, nama4, nama5)

	// Soal 2
	fmt.Println("=========== soal no 2 ==============")
	halo := "Halo Dunia"
	var find = "Dunia"
	var replaceWith = "Golang"
	var newHalo = strings.Replace(halo, find, replaceWith, 1)
	fmt.Println(newHalo)

	// Soal 3
	var kataPertamaSoal3 = "saya"
	var kataKeduaSoal3 = "Senang"
	var kataKetigaSoal3 = "belajar"
	var kataKeempatSoal3 = "golang"

	var find1 = "r"
	var replaceWith1 = "R"

	var newKataKetiga = strings.Replace(kataKetigaSoal3, find1, replaceWith1, 1)

	var find2 = "golang"
	var replaceWith2 = "GOLANG"

	var newKataKeempat = strings.Replace(kataKeempatSoal3, find2, replaceWith2, 2)

	fmt.Println(kataPertamaSoal3, kataKeduaSoal3, newKataKetiga, newKataKeempat)

	//Soal 4
	fmt.Println("=========== soal no 4 ==============")
	var angkaPertama = "8"
	var angkaKedua = "5"
	var angkaKetiga = "6"
	var angkaKeempat = "7"

	newAngka1, _ := strconv.Atoi(angkaPertama)
	newAngka2, _ := strconv.Atoi(angkaKedua)
	newAngka3, _ := strconv.Atoi(angkaKetiga)
	newAngka4, _ := strconv.Atoi(angkaKeempat)

	fmt.Println(newAngka1 + newAngka2 + newAngka3 + newAngka4)

	//Soal 5
	fmt.Println("=========== soal no 5 ==============")
	kalimat := "halo halo bandung"
	angka := 2021

	kalimat = `"` + strings.ReplaceAll(kalimat, "halo", "Hi") + `"` + ` -`

	fmt.Println(kalimat, angka)
}
