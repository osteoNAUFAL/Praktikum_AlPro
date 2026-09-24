package main
import "fmt"

func main()  {
	var a, b int

	fmt.Scan(&a)
	fmt.Scan(&b)

	addition := a + b
	reduction := a - b
	multiplication := a * b
	distribution := a / b
	quotient := a % b

	fmt.Println("Hasil penjumlahan nya adalah ", addition)
	fmt.Println("Hasil pengurangan nya adalah ", reduction)
	fmt.Println("Hasil perkalian nya adalah ", multiplication)
	fmt.Println("Hasil pembagian nya adalah ", distribution)
	fmt.Println("Hasil hasil bagi nya adalah ", quotient)
}