//prosedur

package main
import (
	"fmt"
)

//fungsi rekursif untuk menghitung deret fibonaci
func fibonacci (n int) int {
	if n == 0{
		return 0
	} else if n == 1 {
		return 1
	} else {
		return fibonacci(n-1) + fibonacci (n-2)
	}
}

func main () {
	//menampilkan deret fibonacci hingga suku ke-10
	fmt.Println("Deret fibonacci hingga suku ke-10: ")
	for i := 0; i <= 10; i++ {
		fmt.Printf("Fibonacci (%d) = %d\n",i, fibonacci(i))
	}
}