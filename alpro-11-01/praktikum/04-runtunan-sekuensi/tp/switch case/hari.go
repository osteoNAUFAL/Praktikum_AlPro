package main

import "fmt"

func main() {
    var nomorHari int

    fmt.Print("Masukkan nomor hari (1-7): ")
    fmt.Scanln(&nomorHari)

    switch nomorHari {
    case 1:
        fmt.Println("Hari Senin")
    case 2:
        fmt.Println("Hari Selasa")
    case 3:
        fmt.Println("Hari Rabu")
    case 4:
        fmt.Println("Hari Kamis")
    case 5:
        fmt.Println("Hari Jumat")
    case 6:
        fmt.Println("Hari Sabtu")
    case 7:
        fmt.Println("Hari Minggu")
    default:
        fmt.Println("Nomor hari tidak valid! Masukkan angka 1 sampai 7.")
    }
}