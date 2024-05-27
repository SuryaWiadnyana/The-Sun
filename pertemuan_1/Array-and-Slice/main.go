package main

import "fmt"

func main() {

	// Penerapan Array
	lstAngka := [5]int{0, 1, 2, 3, 4}
	lstAngka1 := [...]int{0, 1, 2, 3, 4}
	lstAngka2 := make([]int, 5)
	lstAngka2[0] = 0
	lstAngka2[1] = 1
	lstAngka2[2] = 2
	lstAngka2[3] = 3
	lstAngka2[4] = 4

	fmt.Println("Angka Pertama 	: 	", lstAngka[0], lstAngka1[0], lstAngka2[0])
	fmt.Println("Angka Kedua 	: 	", lstAngka[1], lstAngka1[1], lstAngka2[1])
	fmt.Println("Angka Ketiga 	: 	", lstAngka[2], lstAngka1[2], lstAngka2[2])

	//Penerapan Array Multi Dimensi
	lstAngka3 := [3][5]int{{0, 1, 2, 3, 4}, {0, 1, 2, 3, 4}, {0, 1, 2, 3, 4}}
	lstAngka4 := [][]int{{0, 1, 2, 3, 4}, {0, 1, 2, 3, 4}, {0, 1, 2, 3, 4}}

	fmt.Println("Angka Keempat 	: 	", lstAngka3)
	fmt.Println("Angka kelima 	: 	", lstAngka4)

	//Pengambilan Data Array
	lstAngka5 := []int{5, 6, 7, 8, 9}

	for i := 0; i < len(lstAngka5); i++ {
		fmt.Println("Angka : ", i+1, "Isinya : ", lstAngka5[i])
	}

	for index, value := range lstAngka5 {
		fmt.Println("Angka ", index+1, ":", value)
	}

	fmt.Println("Angka Keenam : ", lstAngka5[3], "\n")


	//=====================================================

	//Operasi Slice
	fmt.Println("Operasi Slice")

	VarSlice := []int{0, 1, 2, 3, 4, 5, 6}

	fmt.Println("Jumlah Data : ", len(VarSlice))

	VarSlice = append(VarSlice, 12)

	fmt.Println("Isi Data : ", (VarSlice))

	newSlice := make([]int, 3)
	copy(newSlice, VarSlice)

	fmt.Println("Isi Data : ", newSlice)
	
}
