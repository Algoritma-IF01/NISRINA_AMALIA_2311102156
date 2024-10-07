package main

import (
	"fmt"
)
//2311102156- Nisrina Amalia Iffatunnisa
func main() {
	// Deklarasi array untuk menyimpan 5 buah integer
	var integers [5]int

	// Meminta user untuk menginput 5 integer
	fmt.Println("Masukkan 5 angka integer (nilai antara 32 s.d. 127):")
	for i := 0; i < 5; i++ {
		fmt.Scanf("%d", &integers[i])
	}

	// Loop untuk mencetak karakter hasil konversi dari integer berdasarkan ASCII
	fmt.Print("Hasil konversi ke karakter: ")
	for i := 0; i < 5; i++ {
		// Konversi integer ke karakter berdasarkan tabel ASCII
		fmt.Printf("%c", integers[i])
	}
	fmt.Println()

	// Deklarasi string untuk menyimpan input dari user
	var input string

	// Meminta user untuk memasukkan 3 karakter
	fmt.Print("Masukkan 3 karakter: ")

	// Loop untuk memastikan pengguna memasukkan tepat 3 karakter
	for {
		// Meminta user untuk menginput 3 karakter
		fmt.Scanf("%s", &input)

		// Validasi apakah user memasukkan tepat 3 karakter
		if len(input) == 3 {
			break // Keluar dari loop jika input valid
		}
	}

	// Loop untuk mencetak karakter setelah input berdasarkan ASCII
	fmt.Print("Karakter setelah input: ")
	for i := 0; i < 3; i++ {
		// Menambah 1 ke nilai ASCII dari setiap karakter
		fmt.Printf("%c", input[i]+1)
	}
	fmt.Println()
}
