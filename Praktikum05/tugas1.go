package main

import (
	"fmt"
	"math"
)
//2311102156 - Nisrina Amalia Iffatunnisa
// Fungsi untuk menghitung rata-rata
func average(arr []int) float64 {
	sum := 0
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}
	return float64(sum) / float64(len(arr))
}

// Fungsi untuk menghitung standar deviasi
func standardDeviation(arr []int) float64 {
	avg := average(arr)
	var sum float64
	for i := 0; i < len(arr); i++ {
		sum += math.Pow(float64(arr[i])-avg, 2)
	}
	return math.Sqrt(sum / float64(len(arr)))
}

// Fungsi untuk menghitung frekuensi elemen tertentu
func frequency(arr []int, target int) int {
	count := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] == target {
			count++
		}
	}
	return count
}

// Fungsi untuk menghapus elemen pada indeks tertentu
func removeAtIndex(arr []int, index int) []int {
	return append(arr[:index], arr[index+1:]...)
}

func main() {
	var n, x, target, removeIndex int

	// Input jumlah elemen array
	fmt.Print("Masukkan jumlah elemen array: ")
	fmt.Scan(&n)
	arr := make([]int, n)

	// Input elemen array
	fmt.Println("Masukkan elemen array:")
	for i := 0; i < n; i++ {
		fmt.Printf("Elemen %d: ", i)
		fmt.Scan(&arr[i])
	}

	// a. Menampilkan seluruh isi array
	fmt.Println("\na. Menampilkan seluruh isi array:", arr)

	// b. Menampilkan elemen array dengan indeks ganjil
	fmt.Println("b. Menampilkan elemen dengan indeks ganjil:")
	for i := 1; i < len(arr); i += 2 {
		fmt.Printf("%d ", arr[i])
	}
	fmt.Println()

	// c. Menampilkan elemen array dengan indeks genap
	fmt.Println("c. Menampilkan elemen dengan indeks genap:")
	for i := 0; i < len(arr); i += 2 {
		fmt.Printf("%d ", arr[i])
	}
	fmt.Println()

	// d. Menampilkan elemen dengan indeks kelipatan x
	fmt.Print("d. Memasukkan nilai x untuk indeks kelipatan: ")
	fmt.Scan(&x)
	fmt.Printf("Elemen dengan indeks kelipatan %d:\n", x)
	for i := x; i < len(arr); i += x {
		fmt.Printf("%d ", arr[i])
	}
	fmt.Println()

	// e. Menghapus elemen array pada indeks tertentu
	fmt.Print("e. Memasukkan indeks untuk dihapus: ")
	fmt.Scan(&removeIndex)
	arr = removeAtIndex(arr, removeIndex)
	fmt.Println("Array setelah elemen dihapus:", arr)

	// f. Menampilkan rata-rata
	fmt.Printf("f. Rata-rata nilai array: %.2f\n", average(arr))

	// g. Menampilkan standar deviasi
	fmt.Printf("g. Simpangan baku dari array: %.2f\n", standardDeviation(arr))

	// h. Menampilkan frekuensi bilangan tertentu
	fmt.Print("h Memasukkan bilangan untuk menghitung frekuensinya: ")
	fmt.Scan(&target)
	fmt.Printf("Frekuensi %d dalam array: %d\n", target, frequency(arr, target))
}
