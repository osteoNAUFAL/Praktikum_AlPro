package main

import "fmt"

func main()  {
	var nama string
	var skorMatematika, skorBahasaInggris int

	//Membaca input
	fmt.Scan(&nama)
	fmt.Scan(&skorBahasaInggris)
	fmt.Scan(&skorMatematika)

	//Menghitung total dan rata-rata (Pembagian bilangan)
	total := skorBahasaInggris + skorMatematika
	rataRata := total/2

	//Menampilkan output
	fmt.Println(nama)
	fmt.Println(total)
	fmt.Println(rataRata)
}