package main

import (
	"fmt"

	"github.com/federus1105/day1golang/internals"
)

func main() {
	var input float32
	fmt.Println("Masukkan angka")
	fmt.Scanln(&input)
	result := internals.Circle{input}

	wide := result.GetWide()
	around := result.GetAround()

	fmt.Println("Keliling nya adalah", around)
	fmt.Println("Luas nya adalah", wide)
}
