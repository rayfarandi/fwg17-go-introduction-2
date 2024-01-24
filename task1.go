package main

import (
	"math"
)

func Round(x, unit float64) float64 {
	return math.Floor((x/unit)+0.5) * unit
}
