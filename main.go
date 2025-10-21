package main

import (
	"fmt"
	"math"
)

func CalculateArea(r float32) float32 {
	return math.Pi * r * r
}

func Circumference(r float32) float32 {
	return 2 * math.Pi * r
}
func main() {
	// ========= Hitung Luas
	var inputwide float32
	fmt.Println("Masukkan angka")
	fmt.Scanln(&inputwide)
	luas := CalculateArea(inputwide)
	fmt.Printf("Luas lingkaran %2.f\n", luas)

	// ======== Hitung Keliling
	var inputaround float32
	fmt.Println("Masukan angka")
	fmt.Scanln(&inputaround)
	keliling := Circumference(inputaround)
	fmt.Printf("Keliling Lingkaran %2.f\n", keliling)
}
