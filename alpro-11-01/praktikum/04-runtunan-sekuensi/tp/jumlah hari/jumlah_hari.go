package main

import "fmt"

// Fungsi untuk mengecek apakah suatu tahun merupakan tahun kabisat
func isKabisat(tahun int) bool {
	return (tahun%400 == 0) || (tahun%4 == 0 && tahun%100 != 0)
}

func main() {
	var tahun int
	var bulan string

	// Meminta input tahun dan tiga huruf pertama dari nama bulan
	fmt.Scan(&tahun, &bulan)

	// Menentukan jumlah hari berdasarkan input bulan
	switch bulan {
	case "Jan", "Mar", "Mei", "Jul", "Agu", "Okt", "Des":
		fmt.Println(31)
	case "Apr", "Jun", "Sep", "Nov":
		fmt.Println(30)
	case "Feb":
		if isKabisat(tahun) {
			fmt.Println(29)
		} else {
			fmt.Println(28)
		}
	default:
		// Jika input bulan tidak valid)
		fmt.Println("-")
	}
}