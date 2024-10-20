package main

import (
	"fmt"
)

// Definisi fungsi-fungsi matematika
func f(x int) int {
	return x * x // f(x) = x^2
}

func g(x int) int {
	return x - 2 // g(x) = x - 2
}

func h(x int) int {
	return x + 1 // h(x) = x + 1
}

// Fungsi komposisi fogoh = f(g(h(x)))
func fogoh(x int) int {
	return f(g(h(x))) // Komposisi fogoh
}

// Fungsi komposisi gohof = g(h(f(x)))
func gohof(x int) int {
	return g(h(f(x))) // Komposisi gohof
}

// Fungsi komposisi hofog = h(f(g(x)))
func hofog(x int) int {
	return h(f(g(x))) // Komposisi hofog
}

func main() {
	var a, b, c int

	// Input bilangan a, b, c
	fmt.Println("Masukkan tiga angka a, b, c yang dipisahkan oleh spasi:")
	fmt.Scan(&a, &b, &c)

	// Output hasil dari komposisi fungsi dengan format yang diminta
	fmt.Printf("(fogoh) (%d) = %d\n", a, fogoh(a))   // f(g(h(a)))
	fmt.Printf("(gohof) (%d) = %d\n", b, gohof(b))   // g(h(f(b)))
	fmt.Printf("(hofog) (%d) = %d\n", c, hofog(c))   // h(f(g(c)))
}
