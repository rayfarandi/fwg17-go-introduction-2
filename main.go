package main

import (
	"fmt"
)

func main() {
	// task1
	numbers := [...]float64{4.37, 4.324, 4.58, 4.1889, -4.2}
	for _, num := range numbers {
		if num < 0 {
			fmt.Println("Kesalahan:", fmt.Errorf("angka %.2f tidak boleh minus", num))
			continue
		}
		rounded, original, err := Round(num, 0.10)
		if err != nil {
			fmt.Println("Kesalahan:", err)
			continue
		}
		fmt.Printf("%.2f dibulatkan menjadi %.2f\n", original, rounded)
	}

	// task2

	deret := Deret{X: 0, Y: 40, Limit: 40}
	if err := deret.CheckLimit(); err != nil {
		fmt.Println("Error Task 2:", err)
	}
	if err := deret.CheckY(); err != nil {
		fmt.Println("Error Task 2:", err)
	}
	if err := deret.CheckX(); err != nil {
		fmt.Println("Error Task 2:", err)
	} else {
		if prima := deret.AngkaPrima(); len(prima) > 0 {
			fmt.Println("Deret Angka Prima:", prima)
		}
		if ganjil := deret.AngkaGanjil(); len(ganjil) > 0 {
			fmt.Println("Deret Angka Ganjil:", ganjil)
		}
		if genap := deret.AngkaGenap(); len(genap) > 0 {
			fmt.Println("Deret Angka Genap:", genap)
		}
		if fibo := deret.AngkaFibo(); len(fibo) > 0 {
			fmt.Println("Deret Fibonacci:", fibo)
		}
	}

	// // task3
	tabungPerhitungan()

}
