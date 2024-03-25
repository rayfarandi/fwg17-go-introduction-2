package main

import (
	"fmt"
	"math"
)

func Round(x, unit float64) (float64, float64, error) {
	if x < 0 {
		return 0, 0, fmt.Errorf("angka tidak boleh minus")
	}
	rounded := math.Floor((x/unit)+0.5) * unit
	return rounded, x, nil
}
