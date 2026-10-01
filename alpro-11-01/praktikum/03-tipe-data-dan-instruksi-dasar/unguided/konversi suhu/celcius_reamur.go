package main

import "fmt"

func main() {
	var celcius float64

	fmt.Println("Masukkan suhu dalam celcius: ")
	fmt.Scan(&celcius)

	reamur := celcius * 4 / 5

	fmt.Println("Suhu dalam reamur:", reamur)
}