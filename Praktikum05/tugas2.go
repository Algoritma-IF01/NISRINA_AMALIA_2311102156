package main

import "fmt"
//2311102156 Nisrina Amalia Iffatunnisa
func main() {
	var clubA, clubB string
	var scoreA, scoreB int
	results := []string{}
	matchCount := 1

	fmt.Print("Klub A: ")
	fmt.Scan(&clubA)
	fmt.Print("Klub B: ")
	fmt.Scan(&clubB)

	for {
		fmt.Printf("Pertandingan %d: ", matchCount)
		fmt.Scan(&scoreA, &scoreB)

		if scoreA < 0 || scoreB < 0 {
			break
		}

		if scoreA > scoreB {
			results = append(results, clubA)
		} else if scoreB > scoreA {
			results = append(results, clubB)
		} else {
			results = append(results, "Draw")
		}

		matchCount++
	}

	fmt.Println("\nDaftar hasil pertandingan:")
	for i, result := range results {
		if result == "Draw" {
			fmt.Printf("Hasil %d: Draw\n", i+1)
		} else {
			fmt.Printf("Hasil %d: %s\n", i+1, result)
		}
	}

	fmt.Println("Pertandingan selesai")
}
