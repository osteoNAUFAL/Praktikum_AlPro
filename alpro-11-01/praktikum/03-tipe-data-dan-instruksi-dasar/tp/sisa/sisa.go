package main

import "fmt"

func main() {
	var x, y int
	var sisa int

	fmt.Scan(&x, &y)

	sisa = x % y

	fmt.Println("Sisa pembagian:", sisa)
}