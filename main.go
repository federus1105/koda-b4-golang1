package main

import (
	"fmt"
	"math"
)

func CalculateArea(r float32) float32 {
	return math.Pi * r * r
}

func Circumference (r float32) float32 {
	return  2 * math.Pi * r
}
func main() {
	// ========= Hitung Luas
	luas := CalculateArea(10.0)
	fmt.Printf("Luas lingkaran %2.f\n", luas)

	// ======== Hitung Keliling
	keliling := Circumference(10.0)
	fmt.Printf("Keliling Lingkaran %2.f\n", keliling)
}