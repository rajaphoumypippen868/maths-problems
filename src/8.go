package main

import "math/rand"

func RandomInt(min, max int) int {
	return rand.Intn(max - min + 1) + min
}

func RandomFloat64(min, max float64) float64 {
	return rand.Float64()*float64(max-min)+min
}
