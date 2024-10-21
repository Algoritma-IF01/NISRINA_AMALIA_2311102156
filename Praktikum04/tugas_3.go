package main

import (
	"fmt"
	"strings"
)

// Prosedur hitungSkor untuk menghitung total soal yang diselesaikan dan total skor
// Prosedur ini menggunakan parameter formal untuk mengambil input dari pengguna
func hitungSkor_156(nama string, skor *int, soal *int) {
	var waktu int
	// Inisialisasi jumlah soal yang diselesaikan dan total waktu
	*soal = 0
	*skor = 0

	// Membaca 8 waktu penyelesaian dari peserta
	for i := 0; i < 8; i++ {
		fmt.Scan(&waktu) // Membaca waktu untuk setiap soal
		if waktu <= 300 { // Jika waktu penyelesaian kurang dari atau sama dengan 300 menit
			*soal++           // Menambah jumlah soal yang diselesaikan
			*skor += waktu    // Menambahkan waktu ke total skor
		}
	}
}

func main() {
	// Deklarasi variabel untuk menyimpan nama peserta, skor, dan soal
	var nama string
	var totalSkor, totalSoal int
	var pemenang string
	var pemenangSkor, pemenangSoal int = 301, 0 // Inisialisasi pemenang dengan waktu maksimal

	// Meminta input peserta
	for {
		fmt.Println("Masukkan nama peserta dan 8 kali menyelesaikan soal (atau ketik 'selesai' untuk mengakhiri):")
		fmt.Scan(&nama)

		// Memeriksa apakah pengguna ingin mengakhiri input
		if strings.ToLower(nama) == "selesai" {
			break
		}

		// Memanggil prosedur untuk menghitung skor
		hitungSkor_156(nama, &totalSkor, &totalSoal)

		// Mengecek apakah peserta ini menjadi pemenang
		if totalSoal > pemenangSoal || (totalSoal == pemenangSoal && totalSkor < pemenangSkor) {
			pemenang = nama
			pemenangSoal = totalSoal
			pemenangSkor = totalSkor
		}
	}

	// Mencetak hasil pemenang
	fmt.Printf("%s %d %d\n", pemenang, pemenangSoal, pemenangSkor)
}
