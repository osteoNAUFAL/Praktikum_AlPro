package main
import "fmt"

func main(){
	var mil float64
	var kilometer float64

	fmt.Println("Masukkan angka Mil:")
	fmt.Scan(&mil)

	kilometer = mil * 1.6

	fmt.Println("Jarak dalam Kilometer:", kilometer)
}