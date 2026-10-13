package main
import "fmt"

func main(){
	var xi, total int
	fmt.Scan(&xi)
	for xi%2 == 0 {
		total = total + xi
		fmt.Scan(&xi)
	}

	fmt.Println(total)
}




