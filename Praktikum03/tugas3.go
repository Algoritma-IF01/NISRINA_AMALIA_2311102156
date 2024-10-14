package main
//2311102156 Nisrina Amalia Iffatunnisa
import (
	"fmt"
)

func main() {
	var beratParsel_156 int
	var biayaPengiriman_156, biayaTambahan_156, totalBiaya_156 int
	const biayaPerKg_156 = 10000
	const batasGratisTambahan_156 = 10 * 1000 // 10 kg dalam gram

	// Input berat parsel dari pengguna
	fmt.Print("Berat parsel (gram): ")
	fmt.Scan(&beratParsel_156)

	// Menghitung berat total dalam kg dan sisa gram
	beratKg := beratParsel_156 / 1000
	sisaGram := beratParsel_156 % 1000

	// Menghitung biaya pengiriman utama berdasarkan berat kg
	biayaPengiriman_156 = beratKg * biayaPerKg_156

	// Menghitung biaya tambahan berdasarkan sisa gram
	if beratParsel_156 > batasGratisTambahan_156 {
		// Jika berat lebih dari 10kg, sisa gram tidak dikenakan biaya
		biayaTambahan_156 = 0
	} else {
		// Jika berat kurang dari atau sama dengan 10kg
		if sisaGram >= 500 {
			// Jika sisa berat lebih dari atau sama dengan 500 gram
			biayaTambahan_156 = sisaGram * 5
		} else {
			// Jika sisa berat kurang dari 500 gram
			biayaTambahan_156 = sisaGram * 15
		}
	}

	// Menghitung total biaya
	totalBiaya_156 = biayaPengiriman_156 + biayaTambahan_156

	// Menampilkan detail perhitungan
	fmt.Printf("Detail berat: %d kg + %d g\n", beratKg, sisaGram)
	fmt.Printf("Detail biaya: Rp. %d + Rp. %d\n", biayaPengiriman_156, biayaTambahan_156)
	fmt.Printf("Total biaya: Rp. %d\n", totalBiaya_156)
}
