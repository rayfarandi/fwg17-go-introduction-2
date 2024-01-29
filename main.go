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

	deret := Deret{X: 0, Y: 40, Limit: 40}
	prima := deret.AngkaPrima()

	fmt.Println("Deret Angka Prima:", prima)

	ganjil := deret.AngkaGanjil()
	fmt.Println("Deret Angka Ganjil:", ganjil)

	genap := deret.AngkaGenap()
	fmt.Println("Deret Angka Genap:", genap)

	fibo := deret.AngkaFibo()
	fmt.Println("Deret Fibonacci:", fibo)

	// task3
	tabungPerhitungan()

}
