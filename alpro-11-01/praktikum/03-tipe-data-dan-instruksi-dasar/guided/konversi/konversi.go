package main

import "fmt"

func main() {
	var celcius, kelvin float64

	fmt.Print("Masukkan suhu dalam celcius:")
	fmt.Scan(&celcius)

	kelvin = celcius + 273

	fmt.Println("Kelvin: ", kelvin)
}
