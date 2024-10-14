package main
//2311102156 Nisrina Amalia Iffatunnisa
import (
	"fmt"
	"math"
)

// cariFaktor_156 mencari dan mengembalikan semua faktor dari bilangan bulat b.
func cariFaktor_156(b int) []int {
	var faktor []int
	// Mengiterasi dari 1 hingga b untuk mencari faktor
	for i := 1; i <= b; i++ {
		if b%i == 0 { // Memeriksa apakah i adalah faktor dari b
			faktor = append(faktor, i) // Menambahkan faktor ke dalam slice
		}
	}
	return faktor // Mengembalikan daftar faktor
}

// cekPrima_156 memeriksa apakah bilangan bulat b adalah bilangan prima.
func cekPrima_156(b int) bool {
	if b <= 1 {
		return false // Bilangan kurang dari atau sama dengan 1 bukan prima
	}
	// Mengiterasi dari 2 hingga akar kuadrat b untuk memeriksa faktor
	for i := 2; i <= int(math.Sqrt(float64(b))); i++ {
		if b%i == 0 { // Jika b habis dibagi i, berarti bukan prima
			return false
		}
	}
	return true // Jika tidak ada faktor lain, maka b adalah bilangan prima
}

func main() {
	var b int // Deklarasi variabel untuk menyimpan input bilangan

	for {
		fmt.Print("Bilangan: ")

		// Membaca input dari pengguna
		_, err := fmt.Scan(&b) // Menggunakan fmt.Scan untuk membaca bilangan bulat ke dalam variabel b

		// Memeriksa apakah terjadi kesalahan saat membaca input
		if err != nil {
			fmt.Println("Input tidak valid. Silakan masukkan bilangan bulat yang valid.") // Pesan kesalahan jika input tidak valid
			continue // Kembali ke awal loop untuk meminta input lagi
		}

		// Memeriksa apakah b lebih besar dari 1
		if b <= 1 {
			fmt.Println("Masukkan bilangan bulat b > 1.") // Pesan kesalahan jika b tidak lebih dari 1
			continue // Kembali ke awal loop untuk meminta input lagi
		}

		// Mencari faktor menggunakan fungsi cariFaktor_156
		faktor := cariFaktor_156(b)
		fmt.Print("Faktor: ")
		// Menampilkan semua faktor yang ditemukan
		for i, f := range faktor {
			if i == len(faktor)-1 {
				fmt.Println(f) // Menampilkan faktor terakhir dengan newline
			} else {
				fmt.Print(f, " ") // Menampilkan faktor dengan spasi
			}
		}

		// Mengecek apakah bilangan tersebut adalah bilangan prima menggunakan cekPrima_156
		isPrima := cekPrima_156(b)
		fmt.Println("Prima:", isPrima) // Menampilkan hasil cek bilangan prima
		break // Mengakhiri loop setelah mendapatkan input yang valid
	}
}
