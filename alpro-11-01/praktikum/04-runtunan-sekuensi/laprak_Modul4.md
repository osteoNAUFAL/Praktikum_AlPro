# <h1 align="center">Laporan Praktikum Modul 04 - Runtunan dan Sekuensi</h1>
<p align="center">Osteo Naufal Al Badi  - 109092600010</p>

## Dasar Teori

### A. Bahasa Pemrograman Go dan Struktur Dasarnya
Bahasa pemrograman Go (atau Golang) adalah bahasa pemrograman sumber terbuka yang dikembangkan oleh Google untuk membangun perangkat lunak yang sederhana, handal, dan efisien. Menurut Donovan & Kernighan (2015), Go menggabungkan kemudahan penulisan bahasa pemrograman dinamis dengan efisiensi dan keamanan tipe data statis. Setiap program Go dieksekusi melalui fungsi utama main() di dalam package main, di mana pustaka standar seperti fmt digunakan untuk menangani operasi masukan dan keluaran (I/O).

### B. Percabangan (Branching) dan Pengkondisian dalam Go

#### 1. Struktur Percabangan if-else
Percabangan digunakan untuk mengarahkan alur eksekusi program berdasarkan evaluasi kondisi boolean. Jika suatu kondisi bernilai benar (true), blok kode tertentu akan dijalankan; jika bernilai salah (false), program akan melanjutkan ke kondisi alternatif (else if) atau blok alternatif terakhir (else).

#### 2. Struktur Percabangan switch-case
Struktur switch merupakan alternatif dari deretan if-else bertingkat yang membuat kode lebih terstruktur dan mudah dibaca. Dalam bahasa Go, switch dapat ditulis tanpa ekspresi untuk mengevaluasi serangkaian kondisi boolean secara berurutan dari atas ke bawah.

<!-- Tambahkan poin A, B, C, ... atau sub-topik 1, 2, 3, ... sesuai kebutuhan modul -->

## Guided

### 1. grade.go

```go
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var nama string
	var nilai float64
	var grade string

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	nama = scanner.Text()

	fmt.Scan(&nilai)

	if nilai >= 90 && nilai <= 100 {
		grade = "A"
	} else if nilai >= 80 && nilai < 90 {
		grade = "B"
	} else if nilai >= 70 && nilai < 80 {
		grade = "C"
	} else if nilai >= 60 && nilai < 70 {
		grade = "D"
	} else {
		grade = "E"
	}

	fmt.Printf("%s mendapat nilai %s\n", nama, grade)

}

```
#### Deskripsi
Kode di atas merupakan program sederhana dalam bahasa Go yang berfungsi untuk membaca nama dan nilai siswa, lalu menentukan predikat (*grade*) nilainya. Program membaca teks nama siswa (termasuk yang menggunakan spasi) menggunakan `bufio.Scanner`, kemudian mengambil input nilai bertipe desimal menggunakan `fmt.Scan`. Nilai tersebut dievaluasi melalui deretan pengondisian `if-else` untuk menetapkan *grade* dari A hingga E, lalu program menampilkan hasil akhirnya ke layar dengan format nama beserta predikat nilai yang diperoleh.

### 2. penilaian.go

```go
package main

import (
	"bufio" // Digunakan untuk membaca masukan string yang memiliki spasi (seperti nama lengkap)
	"fmt"
	"os" // Menyediakan akses ke sistem operasi, dalam hal ini os.Stdin yang merepresentasikan keyboard
)

func main() {
	// Pendefinisian variabel secara eksplisit beserta tipe datanya
	var pilihan int
	var nama string
	var nilai float64
	var grade string

	// Menampilkan cetakan Menu ke layar
	fmt.Println("============= Menu =============")
	fmt.Println("1. Sistem penilaian")
	fmt.Println("0. Keluar")
	fmt.Print("Pilih opsi: ")

	// Membaca masukan pilihan (angka).
	// Kita menggunakan Scanln agar saat user menekan 'Enter', karakter enter tersebut
	// ikut diolah/dibersihkan dan tidak melompat (mengganggu) masukan nama di bawahnya.
	fmt.Scanln(&pilihan)

	// Percabangan/Sekuensi bersyarat berdasarkan input pengguna
	if pilihan == 1 {
		fmt.Print("Masukkan nama siswa: ")

		// Membuat alat pembaca (scanner) baru yang mengambil masukan dari keyboard
		scanner := bufio.NewScanner(os.Stdin)

		// Proses membaca masukan dari pengguna sampai tombol Enter ditekan
		scanner.Scan()

		// Mengambil teks (nama) yang baru saja dibaca dan menyimpannya ke variabel
		nama = scanner.Text()

		fmt.Print("Masukkan nilai siswa: ")
		fmt.Scanln(&nilai)

		// Menentukan grade nilai
		if nilai >= 90 && nilai <= 100 {
			grade = "A"
		} else if nilai >= 80 && nilai < 90 {
			grade = "B"
		} else if nilai >= 70 && nilai < 80 {
			grade = "C"
		} else if nilai >= 60 && nilai < 70 {
			grade = "D"
		} else {
			grade = "F"
		}

		// %s adalah format (placeholder) untuk mencetak data bertipe string
		// %s pertama akan digantikan oleh isi variabel 'nama', %s kedua oleh 'grade'
		fmt.Printf("%s mendapatkan nilai %s\n", nama, grade)

	} else if pilihan == 0 {
		// Dijalankan jika pengguna mengetik 0
		fmt.Println("Keluar dari program.")

	} else {
		// Dijalankan jika pengguna mengetik angka selain 1 dan 0
		fmt.Println("Pilihan tidak valid.")
	}
}
```
#### Deskripsi
Kode di atas merupakan program CLI berbasis menu interaktif dalam bahasa Go yang berfungsi untuk mengolah dan menentukan huruf mutu (*grade*) nilai siswa. Program pertama-tama menampilkan pilihan menu dan membaca opsi dari pengguna; jika pengguna memilih `1`, program akan meminta masukan nama siswa menggunakan `bufio.Scanner` agar dapat menerima spasi, dilanjutkan dengan masukan nilai bertipe desimal, lalu mengelompokkan nilai tersebut ke dalam *grade* A hingga F melalui deretan pengondisian `if-else` sebelum mencetak hasilnya. Jika pengguna memilih `0`, program akan menampilkan pesan keluar, sedangkan masukan angka lainnya akan menghasilkan pesan bahwa pilihan tidak valid.

### 3. klasifikasi.go

```go
package main

import (
	"fmt" // Hanya memerlukan fmt untuk keperluan input dan output
)

func main() {
	// Pendefinisian variabel secara eksplisit beserta tipe datanya
	var usia int
	var gaji int
	var keterangan string

	// Membaca dua masukan berupa angka dari pengguna (usia dan gaji)
	// fmt.Scan otomatis memisahkan input berdasarkan spasi atau baris baru (enter)
	fmt.Scan(&usia)
	fmt.Scan(&gaji)

	// Menggunakan switch tanpa ekspresi.
	// Cara kerjanya sama persis seperti deretan if - else if.
	// Program akan mengecek dari atas ke bawah, dan menjalankan case pertama yang bernilai benar (true).
	switch {
	case usia < 18:
		keterangan = "Masih sekolah"
	case usia >= 18 && usia <= 25 && gaji >= 50:
		keterangan = "Muda sukses"
	case usia >= 18 && usia <= 25 && gaji < 50:
		keterangan = "Masih belajar hidup"
	case usia >= 26 && usia <= 40 && gaji >= 100:
		keterangan = "Pekerja mapan"
	case usia >= 26 && usia <= 40 && gaji < 100:
		keterangan = "Perlu perbaikan karier"
	case usia > 40 && gaji >= 150:
		keterangan = "Profesional berpengalaman"
	case usia > 40 && gaji < 150:
		keterangan = "Perlu evaluasi finansial"
	}

	// Menampilkan hasil klasifikasi ke layar
	fmt.Println(keterangan)
}
```
#### Deskripsi
Kode di atas merupakan program dalam bahasa Go yang berfungsi untuk mengklasifikasikan status kehidupan dan kondisi finansial seseorang berdasarkan dua nilai masukan, yaitu usia dan jumlah gaji. Setelah membaca input usia dan gaji menggunakan `fmt.Scan`, program memanfaatkan struktur percabangan `switch` tanpa ekspresi yang mengevaluasi rentang usia serta besaran gaji secara bertahap dari atas ke bawah. Hasil evaluasi tersebut disimpan ke dalam variabel `keterangan` dalam bentuk teks (seperti "Masih sekolah", "Muda sukses", hingga "Perlu evaluasi finansial") sebelum akhirnya dicetak ke layar menggunakan `fmt.Println`.

<!-- Tambahkan blok file/kode lain sesuai jumlah file pada soal guided -->

## Unguided

### 1. grade.go (switchcase)

```go
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var nama string
	var nilai float64
	var grade string

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	nama = scanner.Text()

	fmt.Scan(&nilai)

	// Menggunakan switch
	switch {
	case nilai >= 90 && nilai <= 100:
		grade = "A"
	case nilai >= 80 && nilai < 90:
		grade = "B"
	case nilai >= 70 && nilai < 80:
		grade = "C"
	case nilai >= 60 && nilai < 70:
		grade = "D"
	default:
		grade = "E"
	}

	fmt.Printf("%s mendapat nilai %s\n", nama, grade)
}

```

##### Output
<!-- Path relatif dari readme.md ke folder unguided/[nama_soal]/output.png, contoh: -->
![Screenshot Output Unguided](/alpro-11-01/praktikum/04-runtunan-sekuensi/images/grade_switchcase.png)


#### Deskripsi
Pada bagian unguided pertama, program grade.go dimodifikasi dengan mengganti struktur pengondisian if-else menjadi struktur switch tanpa ekspresi. Program membaca nama lengkap siswa menggunakan bufio.Scanner agar dapat menangani spasi, membaca nilai desimal dengan fmt.Scan, mengevaluasi rentang nilai ke dalam predikat huruf mutu (A hingga E) secara bersih menggunakan switch-case, lalu mencetak hasil format nama beserta grade yang diperoleh.

### 2. pajak.go

```go
package main

import "fmt"

func main() {
	var penghasilan, pajak float64

	// Meminta input penghasilan (dalam juta) dari pengguna
	fmt.Scan(&penghasilan)

	// Menggunakan percabangan if-else untuk menghitung pajak sesuai kriteria bracket[cite: 3, 4]
	if penghasilan <= 50 {
		pajak = 0.05 * penghasilan
	} else if penghasilan <= 100 {
		pajak = (0.05 * 50) + (0.10 * (penghasilan - 50))
	} else if penghasilan <= 200 {
		pajak = (0.05 * 50) + (0.10 * 50) + (0.15 * (penghasilan - 100))
	} else {
		pajak = (0.05 * 50) + (0.10 * 50) + (0.15 * 100) + (0.20 * (penghasilan - 200))
	}

	// Mencetak total pajak yang harus dibayar[cite: 4]
	fmt.Println(pajak)
}
```

##### Output
![Screenshot Output Unguided](/alpro-11-01/praktikum/04-runtunan-sekuensi/images/pajak.png)

#### Deskripsi
Pada unguided kedua, dikembangkan program pajak.go untuk menghitung pajak penghasilan progresif berdasarkan bracket pendapatan dalam juta. Program memanfaatkan tipe data float64 untuk menampung nilai numerik presisi tinggi serta menggunakan percabangan bertingkat if-else guna mengkalkulasi besaran pajak secara kumulatif sesuai persentase bertahap (5%, 10%, 15%, dan 20%) sebelum akhirnya menampilkan total pajak yang harus dibayar.

<!-- Duplikasi blok "### [nama_soal]" sesuai jumlah folder soal di dalam unguided -->


## Kesimpulan
Praktikum Modul 04 ini berhasil diselesaikan dengan baik, di mana praktikan memahami dan mampu mengimplementasikan konsep runtunan, sekuensial, serta struktur percabangan (branching) menggunakan if-else dan switch-case dalam bahasa Go. Melalui latihan Guided dan Unguided, praktikan dapat mengelola input data kompleks—seperti pembacaan teks berspasi dengan bufio.Scanner dan angka desimal—serta menerapkannya untuk menyelesaikan permasalahan logika nyata seperti penentuan predikat nilai siswa dan perhitungan pajak penghasilan progresif.

Referensi

## Referensi
1. [Nama Penulis]. ([Tahun]). *[Judul Buku/Sumber]*. [Kota]: [Penerbit]. Diakses pada [tanggal akses] melalui [tautan/DOI]
2. [Nama Penulis]. ([Tahun]). *[Judul Buku/Sumber]*. [Kota]: [Penerbit]. Diakses pada [tanggal akses] melalui [tautan/DOI]
<!-- Tambahkan nomor referensi berikutnya sesuai kebutuhan -->
