package main

import "fmt"

type Deret struct {
	X     int
	Y     int
	Limit int
}

func (d *Deret) AngkaPrima() []int {
	hasil := []int{}
	for i := d.X; i <= d.Y; i++ {
		numPrima := true
		for j := 2; j < i; j++ {
			if i%j == 0 {
				numPrima = false
				break
			}
		}
		if numPrima && i > 1 {
			hasil = append(hasil, i)
		}
	}
	return hasil
}
func (d *Deret) AngkaGanjil() []int {
	hasil := []int{}
	for i := d.X; i <= d.Y; i++ {
		if i%2 != 0 {
			hasil = append(hasil, i)
		}
	}
	return hasil
}

func (d *Deret) AngkaGenap() []int {
	hasil := []int{}
	for i := d.X; i <= d.Y; i++ {
		if i%2 == 0 {
			hasil = append(hasil, i)
		}
	}
	return hasil
}

func (d *Deret) AngkaFibo() []int {
	hasil := []int{0, 1}
	for i := 2; hasil[len(hasil)-1]+hasil[len(hasil)-2] < d.Limit; i++ {
		hasil = append(hasil, hasil[len(hasil)-1]+hasil[len(hasil)-2])
	}
	return hasil
}

func (d *Deret) CheckLimit() error {
	if d.Limit < 0 {
		return fmt.Errorf("limit tidak negatif")
	}
	return nil
}
func (d *Deret) CheckY() error {
	if d.Y < 0 {
		return fmt.Errorf("y tidak negatif")
	}
	return nil
}
func (d *Deret) CheckX() error {
	if d.X < 0 {
		return fmt.Errorf("x tidak negatif")
	}
	return nil
}
