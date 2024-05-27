package main

import (
	"fmt"
)

func main() {

	//============== Condition If...Else =============
	nilai := 80

	if nilai == 100 {
		fmt.Println("Benar")
	} else if nilai < 100 && nilai != 0 {
		fmt.Println("Ada Nilai", "Nilainya:", nilai)
	} else {
		fmt.Println("Tidak Ada nilai ")
	}

	//============== Condition Switch-Case =============
	Nilai := 6
	switch Nilai {
	case 10:
		fmt.Print("nilai sempurna")
	case 9, 8, 7, 6:
		fmt.Print("nilai anda: ", Nilai)
	default:
		fmt.Print("Tidak ada nilai \n")
	}

	//============== Perulangan Pada Golang =============
	for i := 0; i < 5; i++ {
		fmt.Printf("%d \n", i)
	}

	var j = 0

	for j < 5 {
		fmt.Printf("%d \n", j)
		j++
	}

	var k = 0
	for {
		fmt.Printf("%d \n", k)

		k++

		if k > 5 {
			break
		} else {
			continue
		}
	}

	var kv = map[string]int{"a": 0, "b": 1, "c": 2}

	for k, v := range kv {
		fmt.Printf("Key: %s, value: %d \n", k, v)
	}
}
