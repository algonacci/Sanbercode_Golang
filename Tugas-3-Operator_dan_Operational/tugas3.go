package main

import "fmt"

func main() {
	// Soal 1
	fmt.Println("------ SOAL NOMOR 1 ------")
	var panjangPersegiPanjang int = 8
	var lebarPersegiPanjang int = 5

	var alasSegitiga int = 6
	var tinggiSegitiga int = 7

	var kelilingPersegiPanjang = 2 * (panjangPersegiPanjang + lebarPersegiPanjang)
	fmt.Println("Keliling Persegi Panjang adalah", kelilingPersegiPanjang, "cm")

	var luasSegitiga = (alasSegitiga * tinggiSegitiga) / 2
	fmt.Println("Luas Segitiga adalah", luasSegitiga, "cm")

	// Soal 2
	fmt.Println("------ SOAL NOMOR 2 ------")
	var nilaiJohn = 80
	var nilaiDoe = 80
	if nilaiJohn&nilaiDoe >= 80 {
		fmt.Println("Indeksnya A")
	} else if nilaiJohn&nilaiDoe >= 70 {
		fmt.Println("Indeksnya B")
	} else if nilaiJohn&nilaiDoe >= 60 {
		fmt.Println("Indeksnya C")
	} else if nilaiJohn&nilaiDoe >= 50 {
		fmt.Println("Indeksnya D")
	} else {
		fmt.Println("Indeksnya E")
	}

	// Soal 3
	fmt.Println("------ SOAL NOMOR 3 ------")

	var tanggal = "17"
	var bulan = "8"
	var tahun = "1945"
	switch bulan {
	case "1":
		bulan = "Januari"
	case "2":
		bulan = "Februari"
	case "3":
		bulan = "Maret"
	case "4":
		bulan = "April"
	case "5":
		bulan = "Mei"
	case "6":
		bulan = "Juni"
	case "7":
		bulan = "Juli"
	case "8":
		bulan = "Agustus"
	case "9":
		bulan = "September"
	case "10":
		bulan = "Oktober"
	case "11":
		bulan = "November"
	case "12":
		bulan = "Desember"
	}
	fmt.Println(tanggal, bulan, tahun)

	// Soal 4
	fmt.Println("------ SOAL NOMOR 4 ------")
	tahunLahir := 2001

	switch {
	case tahunLahir > 1944 && tahunLahir < 1964:
		fmt.Println("Kamu generasi Baby Boomer")
	case tahunLahir > 1964 && tahunLahir < 1979:
		fmt.Println("Kamu generasi X")
	case tahunLahir > 1979 && tahunLahir < 1993:
		fmt.Println("Kamu generasi Y (Millenials)")
	case tahunLahir > 1994:
		fmt.Println("Kamu generasi Z")
	}
}
