package main

import (
    "fmt"
    "math"
)
//2311102156- Nisrina Amalia Iffatunnisa
func main() {
    // Menggunakan π dari paket math
    const pi = 3.1415926535

    // Menerima input dari pengguna
    var jejari_156 int
    fmt.Print("Program Menghitung Luas Kulit dan Volume pada Bola. ")
    fmt.Print("Jejari = ")
    fmt.Scan(&jejari_156)

    // Menghitung volume dan luas kulit bola
    volume_156 := (4.0 / 3.0) * pi * math.Pow(float64(jejari_156), 3)
    luasKulit_156 := 4 * pi * math.Pow(float64(jejari_156), 2)

    // Menampilkan hasil
    fmt.Printf("Bola dengan jejari %d memiliki volume sebesar %.4f dan luas kulit %.4f\n", jejari_156, volume_156, luasKulit_156)
}
