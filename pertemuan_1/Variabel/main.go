package main

import "fmt"

func main() {

	//Variabel Lebih Dari 1
	var satu, dua, tiga string = "satu", "dua", "tiga"

	empat, lima := "empat", "lima"

	fmt.Printf("%s %s %s \n", satu, dua, tiga)
	fmt.Printf("%s %s \n", empat, lima)

	// Variabel Underscore
	Satu, _ := "satu", "NONE"
	fmt.Printf("%s \n", Satu)

	// Variabel Konstanta

	const phi float32 = 3.14 //Single Variable
	fmt.Printf("%f", phi)

	const (
		Phi float32 = 3.14
		benar bool = true
	)
	fmt.Printf("%f, %t \n", Phi, benar)
}
