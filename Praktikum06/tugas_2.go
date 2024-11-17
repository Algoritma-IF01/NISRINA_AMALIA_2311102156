package main
// 2311102156 Nisrina Amalia Iffatunnisa
import (
	"fmt"
)

func main() {
	var x, y int

	fmt.Print("Masukkan jumlah ikan (x) dan kapasitas wadah (y): ")
	fmt.Scan(&x, &y)

	// array untuk menampung berat ikan
	weights := make([]float64, x)

	fmt.Printf("Masukkan berat %d ikan: ", x)
	for i := 0; i < x; i++ {
		fmt.Scan(&weights[i])
	}

	// menghitung total berat di setiap wadah
	numContainers := (x + y - 1) / y
	containerWeights := make([]float64, numContainers)

	for i := 0; i < x; i++ {
		containerIndex := i / y
		// menambahkan berat ikan ke wadah yang sesuai
		containerWeights[containerIndex] += weights[i]
	}

	// menampilkan total berat ikan di setiap wadah
	fmt.Println("Total berat ikan di setiap wadah:")
	for i := 0; i < len(containerWeights); i++ {
		fmt.Printf("%.2f ", containerWeights[i])
	}
	fmt.Println()

	// menghitung berat rata-rata ikan di setiap wadah
	var totalWeight float64 = 0
	for i := 0; i < len(containerWeights); i++ {
		totalWeight += containerWeights[i]
	}

	averageWeight := totalWeight / float64(numContainers)
	fmt.Printf("Berat rata-rata ikan di setiap wadah: %.2f\n", averageWeight)
}
