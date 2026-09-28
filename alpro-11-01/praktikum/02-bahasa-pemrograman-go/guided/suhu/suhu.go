package main

import "fmt"

func main() {
	// Deklarasi variabel
	var celsius float64

	// Membaca masukan suhu Celsius
	fmt.Scan(&celsius)

	// Menghitung konversi ke Reamur, Fahrenheit, dan Kelvin
	reamur := celsius * 4.0 / 5.0
	fahrenheit := celsius*9.0/5.0 + 32.0
	kelvin := celsius + 273.15

	// Menampilkan keluaran Reamur, Fahrenheit, dan Kelvin
	fmt.Println("Reamur:", reamur)
	fmt.Println("Fahrenheit:", fahrenheit)
	fmt.Println("Kelvin:", kelvin)
}
