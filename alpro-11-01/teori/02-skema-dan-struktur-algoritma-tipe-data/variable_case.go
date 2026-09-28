package main

import "fmt"

func main() {
	var name string

	name = "Osteo Naufal Al Badi"
	fmt.Println("Nama : ", name)

	// name = 17
	// fmt.Println(name)

	var lastName = "Al Badi"
	fmt.Println(lastName)

	middleName := "Naufal"
	fmt.Println("Nama Tengah : ", middleName)

	var (
		fullName = "Osteo Naufal Al Badi"
		firstName = "Osteo"
	)
	
	fmt.Println(fullName)
	fmt.Println(firstName)
}
