package main

import (
	"fmt"
)

func selectionSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIdx] {
				minIdx = j
			}
		}
		arr[i], arr[minIdx] = arr[minIdx], arr[i]
	}
}

func main() {
	var n int
	fmt.Println("MENGURUTKAN NOMOR RUMAH KERABAT HERCULES")
	fmt.Print("Masukkan jumlah wilayah kerabat: ")
	fmt.Scan(&n) 

	if n <= 0 || n >= 1000 {
		fmt.Println("Jumlah wilayah tidak valid!")
		return
	}

	for i := 0; i < n; i++ {
		var m int
		fmt.Printf("Masukkan jumlah nomor rumah kerabat %d: ", i+1)
		fmt.Scan(&m) 

		if m <= 0 || m >= 1000000 {
			fmt.Println("Jumlah nomor rumah tidak valid!")
			return
		}

		rumah := make([]int, m)
		fmt.Printf("Masukkan nomor rumah untuk wilayah %d:\n", i+1)
		for j := 0; j < m; j++ {
			fmt.Scan(&rumah[j])
		}

		selectionSort(rumah)

		fmt.Printf("Hasil pengurutan nomor rumah untuk wilayah %d:\n", i+1)
		for j, num := range rumah {
			if j > 0 {
				fmt.Print(" ")
			}
			fmt.Print(num)
		}
		fmt.Println()
	}
}
