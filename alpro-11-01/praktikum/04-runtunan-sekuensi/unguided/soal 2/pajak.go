package main

import "fmt"

func main() {
	var penghasilan, pajak float64

	// Meminta input penghasilan (dalam juta) dari pengguna
	fmt.Scan(&penghasilan)

	// Menggunakan percabangan if-else untuk menghitung pajak sesuai kriteria bracket[cite: 3, 4]
	if penghasilan <= 50 {
		pajak = 0.05 * penghasilan
	} else if penghasilan <= 100 {
		pajak = (0.05 * 50) + (0.10 * (penghasilan - 50))
	} else if penghasilan <= 200 {
		pajak = (0.05 * 50) + (0.10 * 50) + (0.15 * (penghasilan - 100))
	} else {
		pajak = (0.05 * 50) + (0.10 * 50) + (0.15 * 100) + (0.20 * (penghasilan - 200))
	}

	// Mencetak total pajak yang harus dibayar[cite: 4]
	fmt.Println(pajak)
}