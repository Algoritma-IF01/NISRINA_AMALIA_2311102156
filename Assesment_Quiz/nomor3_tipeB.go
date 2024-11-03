package main

import (
	"fmt"
)
//2311102156 - Nisrina Amalia Iffatunnisa
// nomor 3 jumlah bilangan positif kelipatan 4

// merupakan fungsi rekursif untuk menghitung jumlah bilangan kelipatan 4
func jumlahKelipatan4_156(total int) int {
	var bilangan int
	fmt.Scan(&bilangan)

	// akan menghentikan rekursi jika bilangan negatif
	if bilangan < 0 {
		return total
	}

	// jika bilangan kelipatan 4 dan positif, tambahkan ke total
	if bilangan > 0 && bilangan%4 == 0 {
		total += bilangan
	}

	// Lakukan rekursi dengan total yang sudah diperbarui
	return jumlahKelipatan4_156(total)
}

func main() {
	fmt.Println("Masukkan bilangan (negatif untuk berhenti):")

	// Panggil fungsi rekursif dengan total awal 0
	totalKelipatan4 := jumlahKelipatan4_156(0)

	// Tampilkan hasil total kelipatan 4
	fmt.Printf("Jumlah bilangan kelipatan 4 = %d\n", totalKelipatan4)
}
