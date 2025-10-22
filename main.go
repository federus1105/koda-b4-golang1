package main

import (
	"fmt"
	"math"
)

type Circle struct {
	Radius float32
}

func (c Circle) getWide() float32 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) getAround() float32 {
	return 2 * math.Pi * c.Radius
}

func main() {
	var input float32
	fmt.Println("Masukkan angka")
	fmt.Scanln(&input)
	result := Circle{input}

	wide := result.getWide()
	around := result.getAround()

	fmt.Println("Keliling nya adalah", around)
	fmt.Println("Luas nya adalah", wide)

	// ========= Hitung Luas

	// fmt.Println("Masukkan angka")
	// fmt.Scanln(&inputwide)
	// luas := CalculateArea(inputwide)
	// fmt.Printf("Luas lingkaran %2.f\n", luas)

	// ======== Hitung Keliling
	// var inputaround float32
	// fmt.Println("Masukan angka")
	// fmt.Scanln(&inputaround)
	// keliling := Circumference(inputaround)
	// fmt.Printf("Keliling Lingkaran %2.f\n", keliling)
}
