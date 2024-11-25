package main

import (
	"fmt"
)

func insertionSort(arr []int) {
	n := len(arr)
	for i := 1; i < n; i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}

func checkJarakTetap(arr []int) (bool, int) {
	if len(arr) < 2 {
		return true, 0 
	}

	jarak := arr[1] - arr[0]
	for i := 2; i < len(arr); i++ {
		if arr[i]-arr[i-1] != jarak {
			return false, 0 
		}
	}
	return true, jarak
}

func main() {
	var input int
	var numbers []int

	fmt.Println("Masukkan bilangan bulat (akhiri dengan bilangan negatif):")

	for {
		fmt.Scan(&input)
		if input < 0 {
			break
		}
		numbers = append(numbers, input)
	}

	if len(numbers) == 0 {
		fmt.Println("Tidak ada data untuk diurutkan.")
		return
	}

	insertionSort(numbers)

	fmt.Print("Keluaran = ")
	for i, num := range numbers {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(num)
	}
	fmt.Println()

	isTetap, jarak := checkJarakTetap(numbers)
	if isTetap {
		fmt.Printf("Data berjarak %d\n", jarak)
	} else {
		fmt.Println("Data berjarak tidak tetap")
	}
}
