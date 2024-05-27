package main

import "fmt"

// Numerik Non Desimal
func main() {
	angka1 	:= int8(8)
	var angka2 = 8
	angka3 := -813
	angka4 := -891282881

fmt.Println("ini Hasil Angka 1:  ", angka1)
fmt.Println("ini Hasil Angka 2:  ", angka2)
fmt.Println("ini Hasil Angka 3:  ", angka3)
fmt.Println("ini Hasil Angka 4:  ", angka4)

// Numerik Desimanl
	Angka1 := float32(8)
	Angka2 := 9.54
	Angka3 := 3.83

fmt.Println("ini Hasil Angka 1:  ", Angka1)
fmt.Println("ini Hasil Angka 2:  ", Angka2)
fmt.Println("ini Hasil Angka 3:  ", Angka3)
fmt.Printf("angka 1: %T\n", Angka1)
fmt.Printf("angka 2: %T\n", Angka2)
fmt.Printf("angka 3: %T\n", Angka3)

// Tipe Data String
	var namaDepan string = "Ida Bagus"
	namaBelakang := `"Surya"Wiadnyana"`
	NamaPanggilan := "Surya\"GusSur\""

fmt.Println("Ini Nama Depan: ", namaDepan)
fmt.Println("Ini Nama Belakang: ", namaBelakang)
fmt.Println("Ini Nama Panggilan: ", NamaPanggilan)

// Tipe Data Boolean
	var allowedCondition bool = false
	fmt.Println("Kondisi :", allowedCondition)

	allowedCondition = true
	fmt.Println("Kondisi Terbaru :", allowedCondition)


}