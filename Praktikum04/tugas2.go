package main

import (
	"fmt"
	"math"
)

// Fungsi untuk menghitung jarak antara dua titik (a, b) dan (c, d)
func jarak_156(a, b, c, d float64) float64 {
	return math.Sqrt(math.Pow(a-c, 2) + math.Pow(b-d, 2))
}

// Fungsi untuk menentukan apakah titik (x, y) berada di dalam lingkaran
// dengan titik pusat (cx, cy) dan radius r
func didalam_156(cx, cy, r, x, y float64) bool {
	jarakTitikKePusat_156 := jarak_156(x, y, cx, cy)
	return jarakTitikKePusat_156 <= r
}

// Fungsi utama untuk memproses input dan memberikan output
func main() {
	var cx1_156, cy1_156, r1_156, cx2_156, cy2_156, r2_156, x_156, y_156 float64

	fmt.Println("Masukkan koordinat titik pusat dan radius lingkaran 1:")
	fmt.Scanln(&cx1_156, &cy1_156, &r1_156)
	fmt.Println("Masukkan koordinat titik pusat dan radius lingkaran 2:")
	fmt.Scanln(&cx2_156, &cy2_156, &r2_156)
	fmt.Println("Masukkan koordinat titik sembarang:")
	fmt.Scanln(&x_156, &y_156)

	// Cek posisi titik terhadap kedua lingkaran
	if didalam_156(cx1_156, cy1_156, r1_156, x_156, y_156) && didalam_156(cx2_156, cy2_156, r2_156, x_156, y_156) {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if didalam_156(cx1_156, cy1_156, r1_156, x_156, y_156) {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if didalam_156(cx2_156, cy2_156, r2_156, x_156, y_156) {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}