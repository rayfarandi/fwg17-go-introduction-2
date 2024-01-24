package main

import "fmt"

func DeretAngkaPrima(x, y int) {

	for i := x; i <= y; i++ {
		numPrima := true
		for j := 2; j < i; j++ {

			if i%j == 0 {
				numPrima = false
				break
			}
		}
		if numPrima && i > 1 {

			fmt.Printf("%d ", i)
		}

	}
}
func DeretAngkaGanjil(x, y int) {
	for i := x; i <= y; i++ {
		if i%2 != 0 {
			fmt.Printf("%d ", i)
		}

	}

}

func DeretAngkaGenap(x, y int) {
	for i := x; i <= y; i++ {
		if i%2 == 0 {
			fmt.Printf("%d ", i)
		}
	}
}
