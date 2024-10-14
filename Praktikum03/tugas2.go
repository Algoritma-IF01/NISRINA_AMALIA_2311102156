package main
//2311102156 Nisrina Amalia Iffatunnisa
import (
	"fmt"
)

func hitungAkar2_156(k int) float64 {
	hasil := 1.0

	// Menghitung suku-suku hingga k
	for i := 0; i <= k; i++ {
		numerator := (4*float64(i) + 2) * (4*float64(i) + 2)
		denominator := (4*float64(i) + 1) * (4*float64(i) + 3)
		hasil *= numerator / denominator
	}

	return hasil
}

func main() {
	var k int

	// Meminta input pengguna
	fmt.Print("Nilai K = ")
	fmt.Scan(&k)

	// Menghitung hampiran akar 2
	approxAkar2 := hitungAkar2_156(k)

	// Menampilkan hasil dengan 10 angka di belakang koma
	fmt.Printf("Nilai akar 2 = %.10f\n", approxAkar2)
}
