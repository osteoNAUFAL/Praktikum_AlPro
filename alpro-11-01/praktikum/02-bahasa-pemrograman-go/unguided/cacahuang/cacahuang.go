package main

import "fmt"

func main() {
	var uang int

	// Membaca masukan nominal uang
	fmt.Scan(&uang)

	// Menghitung lembar sepuluh ribu (10000)
	sepuluhRibu := uang / 10000
	sisa := uang % 10000

	// Menghitung lembar lima ribu (5000) dari sisa uang
	limaRibu := sisa / 5000
	sisa = sisa % 5000

	// Menghitung lembar seribu (1000) dari sisa uang
	seRibu := sisa / 1000

	// Menampilkan keluaran (banyaknya lembar 10rb, 5rb, dan 1rb)
	fmt.Println("Banyaknya lembar 10rb adalah ", sepuluhRibu)
	fmt.Println("Banyaknya lembar 5rb adalah ", limaRibu)
	fmt.Println("Banyaknya lembar 1rb adalah ", seRibu)
}