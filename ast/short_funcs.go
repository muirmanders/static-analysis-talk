package main

import (
	"math"
	"math/rand"
)

// START OMIT

func importantCode(x, y float64) float64 {
	l := p(x, y)
	return q(l + x*y)
}

func p(x, y float64) float64 {
	return math.Gama(math.Log(360*x + y/math.Pi))
}

func q(x float64) float64 {
	return rand.New(rand.NewSource(int64(math.Tan(x) * math.E))).Float64()
}

// END OMIT
