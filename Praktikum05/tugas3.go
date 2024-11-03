package main

import (
	"fmt"
)
//2311102156 Nisrina Amalia Iffatunnisa
const NMAX int = 127

type tabel [NMAX]rune

// Fungsi untuk mengisi array hingga karakter TITIK atau jumlah maksimum elemen tercapai
func isiArray_156(t *tabel, n *int) {
	var input rune
	fmt.Println("Masukkan karakter satu per satu (akhiri dengan '.' untuk berhenti):")
	for *n = 0; *n < NMAX; *n++ {
		fmt.Scanf("%c\n", &input)
		if input == '.' {
			break
		}
		(*t)[*n] = input
	}
}

// Fungsi untuk mencetak isi array
func cetakArray(t tabel, n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("%c", t[i])
	}
	fmt.Println()
}

// Fungsi untuk membalik urutan elemen dalam array
func balikanArray(t *tabel, n int) {
	for i := 0; i < n/2; i++ {
		t[i], t[n-1-i] = t[n-1-i], t[i]
	}
}

// Fungsi untuk mengecek apakah array membentuk palindrom
func palindrom(t tabel, n int) bool {
	for i := 0; i < n/2; i++ {
		if t[i] != t[n-1-i] {
			return false
		}
	}
	return true
}

func main() {
	var tab tabel
	var m int

	// Memanggil prosedur untuk mengisi array
	isiArray_156(&tab, &m)

	// Cetak array awal
	fmt.Print("Teks: ")
	cetakArray(tab, m)

	// Memeriksa apakah array membentuk palindrom
	isPalindrome := palindrom(tab, m)
	fmt.Printf("Palindrom? %v\n", isPalindrome)

	// Balikkan array dan cetak
	balikanArray(&tab, m)
	fmt.Print("Reverse teks: ")
	cetakArray(tab, m)
}
