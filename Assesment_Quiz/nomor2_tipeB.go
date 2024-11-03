//2311102156 - Nisrina Amalia Iffatunnisa
// nomor 2 menghitung total biaya makanan di restoran bandung

package main

import "fmt"

func main() {
        var m int
        fmt.Print("Masukkan jumlah rombongan: ")
        fmt.Scan(&m)

        for i := 1; i <= m; i++ {
                var jumlahMenu, jumlahOrang, sisa int
                fmt.Printf("Masukkan jumlah menu, jumlah orang, dan status sisa makanan (0 untuk tidak, 1 untuk iya) :")
                fmt.Scan(&jumlahMenu, &jumlahOrang, &sisa)

                var totalBiaya int
                if jumlahMenu <= 3 {
                        totalBiaya = 10000
                } else if jumlahMenu <= 50 {
                        totalBiaya = 10000 + (jumlahMenu-3) * 2500
                } else {
                        totalBiaya = 100000
                }

                // Tambahan biaya jika makanan sisa
                if sisa == 1 {
                        totalBiaya *= jumlahOrang
                }

                fmt.Printf("Total biaya untuk rombongan %d: Rp %d\n", i, totalBiaya)
        }
}