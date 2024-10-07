package main

import (
    "fmt"
)
//2311102156- Nisrina Amalia Iffatunnisa
func main() {
    var celsius156 float64

    fmt.Print("Masukkan suhu dalam derajat Celsius: ")
    fmt.Scan(&celsius156)

    fahrenheit156 := (celsius156 * 9/5) + 32

    // Menggunakan Println untuk menampilkan hasil tanpa format khusus
    fmt.Println(int(celsius156), "derajat Celsius setara dengan", int(fahrenheit156), "derajat Fahrenheit")
}
