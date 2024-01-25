package main

import (
	"fmt"
	"math"
)

type hitung2d interface {
	luas() float64
	keliling() float64
}
type hitung3d interface {
	volume() float64
}
type hitung interface {
	hitung2d
	hitung3d
}
type tabung struct {
	phi, jari2, tinggi float64
}

func (t tabung) luas() float64 {
	return 2 * t.phi * t.jari2 * (t.jari2 + t.tinggi)
}
func (t tabung) keliling() float64 {
	return 2 * t.phi * t.jari2
}
func (t tabung) volume() float64 {
	return t.phi * math.Pow(t.jari2, 2) * t.tinggi
}
func tabungPerhitungan() {
	t := tabung{jari2: 4, tinggi: 5, phi: 3.14}
	fmt.Println("Perhitungan Luas, Keliling & Volume Tabung")
	fmt.Printf("Data Pengukuran > jari-jari : %.2f , tinggi : %.2f , phi : %.2f\n ", t.jari2, t.tinggi, t.phi)
	fmt.Printf("Luas Tabung: %.2f\n", t.luas())
	fmt.Printf("keliling Tabung: %.2f\n", t.keliling())
	fmt.Printf("Volume Tabung: %.2f\n", t.volume())

}
