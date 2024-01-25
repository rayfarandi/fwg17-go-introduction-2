package main

import (
	"fmt"
)

func main() {

	// task1
	numbers := [...]float64{4.37, 4.32, 4.324, 4.58, 4.1889}

	// Menampilkan hasil
	for _, num := range numbers {
		rounded := Round(num, 0.10)
		fmt.Printf("%.2f dibulatkan menjadi  %.2f\n", num, rounded)
	}

	// task2

	DeretAngkaPrima(0, 40)
	fmt.Println("Deret Angka Prima")

	DeretAngkaGanjil(0, 40)
	fmt.Println("Deret Angka Ganjil")

	DeretAngkaGenap(0, 40)
	fmt.Println("Deret Angka Genap")

	DeretAngkaFibo(40)
	fmt.Println("Deret Angka Fiboonachi")

	// task3
	tabungPerhitungan()

}
