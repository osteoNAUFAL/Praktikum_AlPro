# <h1 align="center">Tugas Pendahuluan Modul 04 - Runtunan/Sekuensi</h1>
<p align="center">Osteo Naufal Al Badi - 109092600010</p>

### 1. Evaluasi Ekspresi Kontrol dalam Go

Diberikan kode dibawah ini:

```go
intNum := 5
intOther := 10
var sngNum float64 = -3
if ____________________ {
 fmt.Println("Beep")
}
```

##### Jawaban:
<!-- Path relatif dari readme.md ke folder unguided/[nama_soal]/output.png, contoh: -->
![Screenshot Output Unguided](/alpro-11-01/praktikum/04-runtunan-sekuensi/images/jawaban%20soal%201.png)

### 2. Tracing: Evaluasi Pernyataan Kondisi

Diberikan kode dibawah ini:

```go
package main
 
 import "fmt"
 
 func main() {
 x := 10
 y := 5
 z := 15
 result := 0
 
 if x > 5 {
 if y < 10 {
 result = x + y
 } else {
 result = x - y
 }
 }
 
 if z > 10 && x == 10 {
 result += z
 } else {
 result = z - x
 }
 
 if x == 10 || y > 10 {
 result += 5
 } else if y == 5 && z > 10 {
 result -= 5
 } else {
 result *= 2
 }
 
 if !(x < 15 && y < 10) {
 result += 10
 } else {
 result -= 10
 }
 
 fmt.Println("Nilai akhir result:", result)
 }
```

#### Pertanyaan
1. Berapa nilai akhir dari variabel result setelah semua pernyataan kondisi dieksekusi?
2. Apa output yang dihasilkan oleh program?
3. Tuliskan langkah-langkah alur eksekusi program berdasarkan kondisi yang diberikan:
 - Kondisi 1: Apakah kondisi x > 5 benar? Jika ya, apa yang terjadi selanjutnya?
-  Kondisi 2: Apakah kondisi z > 10 && x == 10 benar? Bagaimana hal ini memengaruhi nilai result?
-  Kondisi 3: Apakah salah satu dari kondisi x == 10 || y > 10 benar? Apa yang terjadi?
- Kondisi 4: Bagaimana kondisi !(x < 15 && y < 10) dievaluasi? Apa dampaknya pada result?

##### Jawaban
1. 25
2. Nilai akhir result: 25
3. Kondisi 1 (x > 5): Ya, benar ($10 > 5$). Selanjutnya, program masuk ke cabang if di dalamnya dan mengevaluasi y < 10 ($5 < 10$) yang bernilai true, lalu mengeksekusi result = x + y, sehingga nilai result menjadi 15.   
- Kondisi 2 (z > 10 && x == 10): Ya, benar ($15 > 10$ dan $10 == 10$). Hal ini mengakibatkan blok if dieksekusi, yaitu result += z ($15 + 15$), sehingga nilai result bertambah menjadi 30.      
- Kondisi 3 (x == 10 || y > 10): Ya, benar (karena $x == 10$ bernilai true). Yang terjadi adalah program mengeksekusi perintah result += 5 ($30 + 5$), membuat nilai result menjadi 35, serta mengabaikan bagian else if dan else.  
-  Kondisi 4 (!(x < 15 && y < 10)): Ekspresi di dalam kurung bernilai true ($10 < 15$ dan $5 < 10$). Setelah dikenakan negasi !, nilai akhirnya menjadi false. Dampaknya pada result adalah program berpindah ke cabang else dan mengeksekusi result -= 10 ($35 - 10$), sehingga nilai akhir result berkurang menjadi 25.  

### 3. Menentukan Jumlah Hari dalam Sebulan Berdasarkan Tahun dan Bulan

```go
package main

import "fmt"

// Fungsi untuk mengecek apakah suatu tahun merupakan tahun kabisat
func isKabisat(tahun int) bool {
	return (tahun%400 == 0) || (tahun%4 == 0 && tahun%100 != 0)
}

func main() {
	var tahun int
	var bulan string

	// Meminta input tahun dan tiga huruf pertama dari nama bulan
	fmt.Scan(&tahun, &bulan)

	// Menentukan jumlah hari berdasarkan input bulan
	switch bulan {
	case "Jan", "Mar", "Mei", "Jul", "Agu", "Okt", "Des":
		fmt.Println(31)
	case "Apr", "Jun", "Sep", "Nov":
		fmt.Println(30)
	case "Feb":
		if isKabisat(tahun) {
			fmt.Println(29)
		} else {
			fmt.Println(28)
		}
	default:
		// Jika input bulan tidak valid)
		fmt.Println("-")
	}
}
```

##### Output
<!-- Path relatif dari readme.md ke folder unguided/[nama_soal]/output.png, contoh: -->
![Screenshot Output Unguided](/alpro-11-01/praktikum/04-runtunan-sekuensi/images/jumlah_hari.png)


#### Deskripsi
Kode tersebut digunakan untuk menentukan jumlah hari dalam suatu bulan berdasarkan tahun yang dimasukkan. Program menerima input berupa tahun dan tiga huruf pertama nama bulan, kemudian menggunakan `switch` untuk menentukan apakah bulan tersebut memiliki 31 atau 30 hari. Khusus bulan Februari, program menggunakan fungsi `isKabisat()` untuk mengecek apakah tahun tersebut merupakan tahun kabisat, sehingga hasilnya bisa 28 atau 29 hari. Jika nama bulan yang dimasukkan tidak sesuai, program akan menampilkan tanda `-`.

### 4. Switch Case

```go
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
```

##### Output
<!-- Path relatif dari readme.md ke folder unguided/[nama_soal]/output.png, contoh: -->
![Screenshot Output Unguided](/alpro-11-01/praktikum/04-runtunan-sekuensi/images/hari.png)


#### Deskripsi
Kode tersebut digunakan untuk menentukan nama hari berdasarkan nomor hari yang dimasukkan oleh pengguna. Program meminta input angka 1 sampai 7, kemudian menggunakan `switch` untuk mencocokkan angka dengan nama hari, yaitu 1 untuk Senin hingga 7 untuk Minggu. Jika pengguna memasukkan angka selain 1–7, program akan menampilkan pesan bahwa nomor hari tidak valid.

## Kesimpulan
 Berdasarkan praktikum yang telah dilakukan, dapat disimpulkan bahwa struktur kontrol dalam bahasa Go seperti `if`, `else`, dan `switch` dapat digunakan untuk membuat program mengambil keputusan berdasarkan kondisi tertentu. Selain itu, penggunaan operator logika dan fungsi juga membantu dalam menentukan hasil dari suatu kondisi, seperti pengecekan tahun kabisat dan penentuan nama hari. Dengan memahami struktur kontrol tersebut, program dapat dibuat untuk menghasilkan output yang sesuai berdasarkan input dan kondisi yang diberikan.