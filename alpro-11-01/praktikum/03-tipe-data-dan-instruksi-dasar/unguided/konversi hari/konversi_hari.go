package main

import "fmt"

func main() {
	var totalHari int
	fmt.Println("Masukkan jumlah hari: ")
	fmt.Scan(&totalHari)

	tahun := totalHari / 360
	sisaHari := totalHari % 360

	bulan := sisaHari / 30
	sisaHari = sisaHari % 30

	minggu := sisaHari / 7
	sisaHari = sisaHari % 7

	fmt.Println("Tahun:", tahun)
	fmt.Println("Bulan:", bulan)
	fmt.Println("Minggu:", minggu)
	fmt.Println("Sisa Hari:", sisaHari)
}
