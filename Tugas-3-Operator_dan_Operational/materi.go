package main

import (
	"fmt"
)

func main() {
	// operator penjumlahan
	jumlah := 8 + 3
	fmt.Println(jumlah) // hasilnya 13

	// operator pengurangan
	kurang := 8 - 3
	fmt.Println(kurang) // hasilnya 5

	// operator perkalian
	kali := 8 * 3
	fmt.Println(kali) // hasilnya 24

	// operator pembagian
	bagi := 8 / 4
	fmt.Println(bagi) // hasilnya 2

	// operator modulus
	modulus := 8 % 3
	fmt.Println(modulus) // hasilnya 2

	var angka = 8
	fmt.Println(angka) // 8
	angka += 10
	fmt.Println(angka) // 18

	var angka2 = 5
	fmt.Println(angka2) // 5
	angka2 += 5
	fmt.Println(angka2) // 10

	var angkaLagi = 8

	fmt.Println(angkaLagi > 5) // true

	fmt.Println(angkaLagi < 5) // false

	fmt.Println(angkaLagi >= 5) // true

	fmt.Println(angkaLagi <= 5) // false

	fmt.Println(angkaLagi == 5) // false

	fmt.Println(angkaLagi != 5) // true

	var a = true
	var b = false
	var c = true
	var d = false

	fmt.Println(a && c) // true

	fmt.Println(a && b) // false

	fmt.Println(a || b) // true

	fmt.Println(b || d) // false

	fmt.Println(!b && !d) // true

	fmt.Println(!a || b) // false

	var mood = "happy"
	if mood == "happy" {
		fmt.Println("hari ini aku bahagia!")
	}

	var minimarketStatus = "open"
	var telur = "soldout"
	var buah = "soldout"
	if minimarketStatus == "open" {
		fmt.Println("saya akan membeli telur dan buah")
		if telur == "soldout" || buah == "soldout" {
			fmt.Println("belanjaan saya tidak lengkap")
		} else if telur == "soldout" {
			fmt.Println("telur habis")
		} else if buah == "soldout" {
			fmt.Println("buah habis")
		}
	} else {
		fmt.Println("minimarket tutup, saya pulang lagi")
	}

	var buttonPushed = 1
	switch buttonPushed {
	case 1:
		fmt.Println("matikan TV!")
	case 2:
		fmt.Println("turunkan volume TV!")
	case 3:
		fmt.Println("tingkatkan volume TV!")
	case 4:
		fmt.Println("matikan suara TV!")
	default:
		fmt.Println("Tidak terjadi apa-apa")
	}

	var point = 6

	switch {
	case point == 8:
		fmt.Println("perfect")
	case (point < 8) && (point > 3):
		fmt.Println("awesome")
		fallthrough
	case point < 5:
		fmt.Println("you need to learn more")
	default:
		{
			fmt.Println("not bad")
			fmt.Println("you need to learn more")
		}
	}
}
