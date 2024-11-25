package main

import (
	"fmt"
)

type Buku struct {
	id        string
	judul     string
	penulis   string
	penerbit  string
	eksemplar int
	tahun     int
	rating    int
}

const nMax = 7919

type DaftarBuku [nMax]Buku

func DaftarkanBuku(pustaka *DaftarBuku, n *int) {
	fmt.Println("Masukkan data buku:")
	for i := 0; i < *n; i++ {
		var buku Buku
		fmt.Printf("Data buku ke-%d:\n", i+1)
		fmt.Print("ID: ")
		fmt.Scanln(&buku.id)

		fmt.Print("Judul: ")
		fmt.Scanln(&buku.judul)

		fmt.Print("Penulis: ")
		fmt.Scanln(&buku.penulis)

		fmt.Print("Penerbit: ")
		fmt.Scanln(&buku.penerbit)

		fmt.Print("Eksemplar: ")
		fmt.Scanln(&buku.eksemplar)

		fmt.Print("Tahun: ")
		fmt.Scanln(&buku.tahun)

		fmt.Print("Rating: ")
		fmt.Scanln(&buku.rating)

		pustaka[i] = buku
		fmt.Println()
	}
}

// Procedure Cetak Terfavorit
func CetakTerfavorit(pustaka *DaftarBuku, n int) {
	if n == 0 {
		fmt.Println("Tidak ada buku dalam pustaka.")
		return
	}
	maxRating := pustaka[0]
	for i := 1; i < n; i++ {
		if pustaka[i].rating > maxRating.rating {
			maxRating = pustaka[i]
		}
	}
	fmt.Printf("%s\n", maxRating.judul) 
}

// Procedure Urut Buku dengan Insertion Sort
func UrutBuku(pustaka *DaftarBuku, n int) {
	for i := 1; i < n; i++ {
		key := pustaka[i]
		j := i - 1
		for j >= 0 && pustaka[j].rating < key.rating { 
			pustaka[j+1] = pustaka[j]
			j--
		}
		pustaka[j+1] = key
	}
}

func Cetak5Terbaru(pustaka *DaftarBuku, n int) {
	// Menampilkan lima buku dengan rating tertinggi dengan nomor urut
	for i := 0; i < n && i < 5; i++ {
		fmt.Printf("%d. %s\n", i+1, pustaka[i].judul)
	}
}

func CariBuku(pustaka *DaftarBuku, n int, r int) {
	left, right := 0, n-1
	for left <= right {
		mid := (left + right) / 2
		if pustaka[mid].rating == r {
			buku := pustaka[mid]
			fmt.Printf("%s,%s,%s,%d,%d,%d\n",
				buku.judul, buku.penulis, buku.penerbit, buku.tahun, buku.eksemplar, buku.rating)
			return
		} else if pustaka[mid].rating > r {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	fmt.Println("Tidak ditemukan") 
}

func main() {
	var pustaka DaftarBuku
	var n, ratingCari int

	fmt.Print("Masukkan jumlah buku: ")
	fmt.Scanln(&n)

	if n <= 0 || n > nMax {
		fmt.Println("Jumlah buku tidak valid.")
		return
	}

	DaftarkanBuku(&pustaka, &n)

	fmt.Println(" Buku terfavorit):")
	CetakTerfavorit(&pustaka, n)

	UrutBuku(&pustaka, n)

	fmt.Println(" Lima buku dengan rating tertinggi:") 
	Cetak5Terbaru(&pustaka, n)

	fmt.Print("Masukkan rating buku yang akan dicari: ")
	fmt.Scanln(&ratingCari)
	CariBuku(&pustaka, n, ratingCari)
}
