package main

import "fmt"

func cetak2(ch chan int) {
	fmt.Println("ini dari goroutine...")
	ch <- 10
}

func main2() {
	angka := make(chan int)
	go cetak(angka)
	nilai := <-angka
	fmt.Println("nilai channel integer :", nilai)
	fmt.Println("ini dari function main...")
}
