package main

import "fmt"

func main() {

	varMaps := map[string]string{
		"Peserta_1": "Surya",
		"Peserta_2": "Bagus",
	}

	fmt.Println("Map Sekarang :", varMaps)

	delete(varMaps, "Peserta_1")

	fmt.Println("Map sekarang :", varMaps)

	value, isExist := varMaps["Peserta_1"]

	fmt.Println("Key Peserta_1 Ada :", isExist, "Value :", value)

	value, isExist = varMaps["Peserta_2"]

	fmt.Println("Key Peserta_2 Ada :", isExist, "Value :", value)

	//Slice Map
	VarSliceMaps := []map[string]string{
		{"nama": "Surya", "Asal": "SMP Dwijendra"},
		{"nama": "Bagus", "Asal": "SMA 2 Denpasar"},
	}

	fmt.Println("Nama Peserta 1 :", VarSliceMaps[0]["nama"])

	for index, DataPeserta := range VarSliceMaps {
		fmt.Printf("Peserta %d : %s \n", index+1, DataPeserta["nama"])
	}
}
