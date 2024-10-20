package main

import (
	"fmt"
)

// Prosedur cetakDeret_156 untuk mencetak deret Skiena
// Prosedur tidak mengembalikan nilai, hanya mencetak deret
func cetakDeret_156(n_156 int) {
	// Mencetak setiap suku dari deret
	for n_156 != 1 {
		fmt.Print(n_156, " ") // Cetak nilai suku saat ini
		if n_156%2 == 0 {
			n_156 /= 2 // Jika genap, bagi 2
		} else {
			n_156 = 3*n_156 + 1 // Jika ganjil, kalikan 3 dan tambah 1
		}
	}
	// Cetak nilai terakhir yaitu 1
	fmt.Print(n_156, "\n")
}

func main() {
	// Deklarasi variabel untuk input
	var n_156 int

	// Meminta input dari pengguna
	fmt.Println("Masukkan bilangan bulat positif (kurang dari 1 juta):")
	fmt.Scan(&n_156)

	// Memanggil prosedur untuk mencetak deret
	cetakDeret_156(n_156)
}
