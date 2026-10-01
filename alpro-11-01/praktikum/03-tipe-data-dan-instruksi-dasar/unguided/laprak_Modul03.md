# <h1 align="center">Laporan Praktikum Modul 3 - Variabel dan Operator</h1>
<p align="center">Osteo Naufal Al Badi - 109092600010</p>

## Dasar Teori

### A. Variabel dan Tipe Data dalam Bahasa Go
Variabel merupakan wadah dalam memori komputer yang digunakan untuk menyimpan suatu nilai dengan tipe data tertentu selama program berjalan. Bahasa Go (Golang) merupakan bahasa berjenis *statically typed*, yang berarti setiap variabel wajib memiliki tipe data yang jelas saat dideklarasikan. Tipe data dasar pada bahasa Go meliputi tipe numerik bulat (`int`), tipe numerik pecahan (`float64`), tipe nilai kebenaran (`bool`), serta tipe teks (`string`).

### B. Operator Aritmatika dan Pengisian Nilai

#### 1. Operator Aritmatika
Operator aritmatika digunakan untuk melakukan operasi matematika dasar pada tipe data numerik. Beberapa operator utama dalam Go meliputi penjumlahan (`+`), pengurangan (`-`), perkalian (`*`), pembagian (`/`), dan sisa hasil bagi atau modulus (`%`). Operasi modulus sangat berguna untuk memisahkan digit angka, menentukan kelipatan, atau mengonversi satuan.

#### 2. Operator Penugasan dan Deklarasi Singkat
Pengisian nilai ke dalam variabel dilakukan menggunakan operator penugasan (`=`). Selain itu, Go menyediakan sintaks deklarasi singkat (*short variable declaration*) menggunakan operator `:=` yang secara otomatis menentukan tipe data variabel berdasarkan nilai masukan (*type inference*).

<!-- Tambahkan poin A, B, C, ... atau sub-topik 1, 2, 3, ... sesuai kebutuhan modul -->

## Guided

### 1. konversi.go

```go
package main

import "fmt"

func main() {
	var celcius, kelvin float64

	fmt.Print("Masukkan suhu dalam celcius:")
	fmt.Scan(&celcius)

	kelvin = celcius + 273

	fmt.Println("Kelvin: ", kelvin)
}
```
#### Deskripsi
Program ini bertujuan untuk mengonversi nilai suhu dari satuan Celcius ke Kelvin. Program menerima masukan nilai Celcius bertipe float64 dari pengguna melalui fmt.Scan(), kemudian menambahkan konstanta 273 pada nilai Celcius tersebut untuk menghasilkan nilai suhu dalam satuan Kelvin, lalu menampilkan hasilnya ke layar.

### 2. tukar.go

```go
package main

import "fmt"

func main() {
	var x, y, z int
	fmt.Scan(&x, &y, &z)

	temp := x
	x = y
	z = y
	y = temp

	fmt.Println(x, y, z)
}
```
#### Deskripsi
Program ini digunakan untuk melakukan penukaran dan penyalinan nilai antar variabel integer (x, y, z). Program menyimpan nilai awal x ke dalam variabel sementara temp, kemudian mengganti nilai x dengan nilai y, mengisi nilai z dengan nilai y, dan mengganti nilai y dengan nilai temp. Hasil penukaran posisi nilai variabel kemudian dicetak secara berurutan.

### 3. kasir.go

```go
package main

import "fmt"

func main() {
	var x int
	fmt.Scan(&x)

	var sepuluhRibuan int = x / 10000
	var sisa int = x % 10000

	var limaRibuan int = sisa / 5000
	sisa = sisa % 5000

	var seRibuan int = sisa / 1000

	fmt.Println("Sepuluh ribuan:", sepuluhRibuan)
	fmt.Println("Lima ribuan:", limaRibuan)
	fmt.Println("Seribuan:", seRibuan)
}
```
#### Deskripsi
Program ini berfungsi untuk memecah sejumlah uang pecahan kelipatan ribuan menjadi lembaran pecahan 10.000, 5.000, dan 1.000. Program menggunakan operasi pembagian integer (/) untuk menentukan jumlah lembar dari masing-masing pecahan dan operator modulus (%) untuk mendapatkan sisa uang yang belum dipecahkan pada setiap tahap.

<!-- Tambahkan blok file/kode lain sesuai jumlah file pada soal guided -->

## Unguided

### 1. celcius_reamur.go

```go
package main

import "fmt"

func main() {
	var celcius float64

	fmt.Println("Masukkan suhu dalam celcius: ")
	fmt.Scan(&celcius)

	reamur := celcius * 4 / 5

	fmt.Println("Suhu dalam reamur:", reamur)
}
```

##### Output
<!-- Path relatif dari readme.md ke folder unguided/[nama_soal]/output.png, contoh: -->
![Screenshot Output Unguided](/alpro-11-01/praktikum/03-tipe-data-dan-instruksi-dasar/unguided/images/celcius%20to%20reamur.png)


#### Deskripsi
Program ini mengimplementasikan rumus konversi suhu dari Celcius ke Reamur. Pengguna memasukkan nilai suhu Celcius dalam tipe float64, lalu program menghitung hasilnya menggunakan rumus reamur = celcius * 4 / 5 dan menampilkan hasil konversinya ke layar.

### 2. konversi_hari.go

```go
package main

import "fmt"

func main() {
	var totalHari int
	fmt.Scan(&totalHari)

	tahun := totalHari / 360
	sisaHari := totalHari % 360

	bulan := sisaHari / 30
	sisaHari = sisaHari % 30

	minggu := sisaHari / 7
	sisaHari = sisaHari % 7

	fmt.Println("Tahun:", tahun)
	fmt.Println("Bulan:", bulan)
	fmt.Println("Minggu:", minggu)
	fmt.Println("Sisa Hari:", sisaHari)
}
```

##### Output
![Screenshot Output Unguided](/alpro-11-01/praktikum/03-tipe-data-dan-instruksi-dasar/unguided/images/Screenshot%202026-10-01%20112255.png)

#### Deskripsi
Program ini mengonversi masukan total jumlah hari ke dalam satuan tahun, bulan, minggu, dan sisa hari berdasarkan asumsi standar waktu (1 tahun = 360 hari, 1 bulan = 30 hari, 1 minggu = 7 hari). Program secara bertahap menghitung kuantitas setiap satuan menggunakan pembagian bulat (/) dan memperbarui sisa hari menggunakan operator modulus (%).

<!-- Duplikasi blok "### [nama_soal]" sesuai jumlah folder soal di dalam unguided -->


## Kesimpulan
Praktikum Modul 3 ini memberikan pemahaman mendasar mengenai deklarasi variabel, pemilihan tipe data yang sesuai (int, float64), serta pemanfaatan operator aritmatika dalam bahasa pemrograman Go. Melalui latihan program Guided dan Unguided, mahasiswa berhasil mengimplementasikan operasi matematis seperti konversi suhu, penukaran variabel, manipulasi nilai pecahan mata uang, serta konversi waktu menggunakan kombinasi operator pembagian dan modulus.

## Referensi
1. [Nama Penulis]. ([Tahun]). *[Judul Buku/Sumber]*. [Kota]: [Penerbit]. Diakses pada [tanggal akses] melalui [tautan/DOI]
2. [Nama Penulis]. ([Tahun]). *[Judul Buku/Sumber]*. [Kota]: [Penerbit]. Diakses pada [tanggal akses] melalui [tautan/DOI]
<!-- Tambahkan nomor referensi berikutnya sesuai kebutuhan -->
