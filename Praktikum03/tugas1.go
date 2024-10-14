package main
//2311102156 Nisrina Amalia Iffatunnisa
import (
	"fmt"
	"math"
)

func main() {
	var kantong1, kantong2 float64

	for {
		// Meminta input berat barang di kedua kantong
		fmt.Print("Masukan berat belanjaan di kedua kantong: ")
		_, err := fmt.Scan(&kantong1, &kantong2)
		if err != nil {
			fmt.Println("Input tidak valid, coba lagi.")
			continue
		}

		// Mengecek apakah ada kantong dengan berat negatif
		if kantong1 < 0 || kantong2 < 0 {
			fmt.Println("Proses selesai.")
			break
		}

		// Mengecek apakah total berat kedua kantong melebihi 150 kg
		totalBerat_156 := kantong1 + kantong2
		if totalBerat_156 > 150 {
			fmt.Println("Proses selesai.")
			break
		}

		// Menghitung selisih berat antara kedua kantong
		selisihBerat_156 := math.Abs(kantong1 - kantong2)

		// Mengecek apakah selisih berat lebih dari atau sama dengan 9 kg
		if selisihBerat_156 >= 9 {
			fmt.Println("Sepeda motor pak Andi akan oleng: true")
		} else {
			fmt.Println("Sepeda motor pak Andi akan oleng: false")
		}
	}
}
