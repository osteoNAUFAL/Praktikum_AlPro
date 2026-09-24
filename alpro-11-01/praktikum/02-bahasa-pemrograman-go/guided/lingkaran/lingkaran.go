package main

import "fmt"

func main() {
	// Deklarasi variabel
	var r, luas float64

	// Assignment nilai pi
	pi := 3.14

	// Membaca masukan jari-jari lingkaran
	fmt.Scan(&r)

	// Menghitung luas lingkaran
	luas = pi * r * r

	// Menampilkan hasil luas lingkaran
	fmt.Println("Luas lingkaran:", luas)
}
