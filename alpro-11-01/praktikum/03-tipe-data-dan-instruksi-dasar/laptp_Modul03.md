# <h1 align="center">Tugas Pendahuluan Modul 3 - Variabel dan Operator</h1>
<p align="center">Osteo Naufal Al Badi - 109092600010</p>

### 1. sisa.go

```go
package main

import "fmt"

func main() {
	var x, y int
	var sisa int

	fmt.Scan(&x, &y)

	sisa = x % y

	fmt.Println("Sisa pembagian:", sisa)
}
```

##### Output
<!-- Path relatif dari readme.md ke folder unguided/[nama_soal]/output.png, contoh: -->
![Screenshot Output](../../../../screenshots/sisa.png)



#### Deskripsi

 Rincian Alur Kerja Program:

1. **Deklarasi Variabel:** Mendeklarasikan variabel `x`, `y`, dan `sisa` dengan tipe data integer (`int`).
2. **Input Data:** Membaca dua nilai masukan dari pengguna menggunakan `fmt.Scan(&x, &y)`.
3. **Proses Operasi:** Menghitung sisa hasil bagi dengan rumus `sisa = x % y`.
4. **Output Data:** Menampilkan hasil akhir berupa sisa pembagian menggunakan `fmt.Println()`.

---

### 2. konversi.go

```go
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
```

##### Output
<!-- Path relatif dari readme.md ke folder unguided/[nama_soal]/output.png, contoh: -->
![Screenshot Output](../../../screenshots/konversi.png)


#### Deskripsi

 Rincian Alur Kerja Program:

1. **Deklarasi Variabel:** Mendeklarasikan variabel `mil` dan `kilometer` bertipe data pecahan/desimal (`float64`).
2. **Input Data:** Menampilkan instruksi input dan membaca nilai mil yang dimasukkan pengguna melalui `fmt.Scan(&mil)`.
3. **Proses Konversi:** Menghitung konversi jarak dengan rumus `kilometer = mil * 1.6`.
4. **Output Data:** Menampilkan hasil akhir jarak dalam satuan kilometer menggunakan `fmt.Println()`.

---

### 3. bool.go

```go
package main

import "fmt"

func main() {
	var nilai bool

	fmt.Scan(&nilai)

	fmt.Println(nilai)
}
```

##### Output
<!-- Path relatif dari readme.md ke folder unguided/[nama_soal]/output.png, contoh: -->
![Screenshot Output](../../../../screenshots/bool.png)


#### Deskripsi
Rincian Alur Kerja Program:

1. **Deklarasi Variabel:** Mendeklarasikan variabel `nilai` dengan tipe data boolean (`bool`).
2. **Input Data:** Membaca nilai masukan boolean (`true` atau `false`) dari pengguna menggunakan `fmt.Scan(&nilai)`.
3. **Output Data:** Menampilkan kembali nilai boolean yang telah disimpan ke layar menggunakan `fmt.Println(nilai)`.


## Kesimpulan
Ketiga program di atas merupakan penerapan dasar dari variabel, tipe data, serta operasi masukan (input) dan keluaran (output) menggunakan bahasa Go:

1. **Program Sisa Pembagian (Modulus):** Mengolah tipe data bilangan bulat (`int`) untuk menghitung dan menampilkan sisa hasil bagi dari dua nilai menggunakan operator `%`.
2. **Program Konversi Mil ke Kilometer:** Mengolah tipe data pecahan/desimal (`float64`) untuk melakukan perhitungan aritmatika konversi satuan jarak dengan perkalian faktor `1.6`.
3. **Program I/O Boolean:** Mengolah tipe data logika (`bool`) untuk membaca input nilai kebenaran (`true`/`false`) dan mencetaknya kembali secara langsung.

Secara keseluruhan, praktikum ini memperlihatkan bagaimana fungsi `fmt.Scan()` digunakan untuk menerima berbagai jenis tipe data dari pengguna dan fungsi `fmt.Println()` untuk menampilkan hasilnya.