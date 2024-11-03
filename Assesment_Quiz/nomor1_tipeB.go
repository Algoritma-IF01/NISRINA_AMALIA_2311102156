//2311102156 - Nisrina Amalia Iffatunnisa
// nomor 1 - mengecek voucher untuk mahasiswa Universitas Telkom

package main

import "fmt"

func main() {
	var n_156 int // input jumlah mahasiswa Telkom
	fmt.Print("Masukkan jumlah mahasiswa Telkom: ")
	fmt.Scan(&n_156) 

	validCount := 0
	invalidCount := 0

	for i := 0; i < n_156; i++ {
		var serial string
		fmt.Print("Masukkan kode voucher mahasiswa :  ")
		fmt.Scan(&serial) // input nomor seri voucher

		// mengecek panjang nomor seri (harus 5 atau 6)
		length := len(serial)
		if length != 5 && length != 6 {
			invalidCount++
			continue
		}

		// mengambil dua digit pertama dan dua digit terakhir dari string sebagai int
		first := int(serial[0] - '0')
		second := int(serial[1] - '0')
		secondLast := int(serial[length-2] - '0')
		last := int(serial[length-1] - '0')

		// mengecek apakah perkalian dua digit pertama dan dua digit terakhir sama
		if first*second != secondLast*last {
			invalidCount++
			continue
		}

		// mengecek apakah digit tengah (atau tiga digit tengah) genap
		if length == 5 {
			middle := int(serial[2] - '0') // digit tengah untuk panjang 5
			if middle%2 != 0 {
				invalidCount++
				continue
			}
		} else if length == 6 {
			// digit ketiga, keempat, dan kelima untuk panjang 6
			midFirst := int(serial[2] - '0')
			midSecond := int(serial[3] - '0')
			midThird := int(serial[4] - '0')

			if midFirst%2 != 0 || midSecond%2 != 0 || midThird%2 != 0 {
				invalidCount++
				continue
			}
		}

		// jika semua kondisi di atas terpenuhi, maka voucher valid
		validCount++
	}

	// output hasil
	fmt.Println(validCount, invalidCount) //akan menampilkan yang valid dulu, baru yang tidak valid
}
