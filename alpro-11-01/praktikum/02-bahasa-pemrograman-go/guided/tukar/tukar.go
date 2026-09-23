package main

import "fmt"

func main()  {
	var angkaA, angkaB int

	fmt.Scan(&angkaA)
	fmt.Scan(&angkaB)

	angkaA,angkaB = angkaB,angkaA

	fmt.Println(angkaA)
	fmt.Println(angkaB)
}