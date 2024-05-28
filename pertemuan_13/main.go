package main

import (
	"fmt"
	"pertemuan_13/config"
	"pertemuan_13/funcs"

	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println(err)
	}

	// Get token from config
	token := config.GetToken()
	urlAPI := "http://127.0.0.1:8000/buku/"
	id := "664da4c816bb6799faf5da16"
	

	// Get All
	data, err := funcs.GetAll(urlAPI, token)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(data)
	}
	fmt.Println("=============================================")

	// Get Buku By ID
	datas, err := funcs.GetBukuByID(urlAPI, id, token)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(datas)
	}
	fmt.Println("=============================================")

	// Create Buku
	bdy := map[string]interface{}{
		"judul":   "Belajar Golang2",
		"penulis": "Surya",
		"halaman": 200,
		"status":  "tersedia",
	}

	err = funcs.CreateBuku(urlAPI, bdy, token)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Buku Berhasil Dibuat")
	}
	fmt.Println("=============================================")

	// Delete Buku
	err = funcs.DeleteBuku(urlAPI, id, token)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Delete Buku Berhasil")
	}

	fmt.Println("=============================================")

	// Update Buku
	UpdBody := map[string]interface{}{
		"judul":   "THE SUN",
		"penulis": "Surya",
		"halaman": 100,
		"status":  "tersedia",
	}

	err = funcs.UpdateBuku(urlAPI, id, UpdBody, token)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Update Buku Berhasil")
	}
}
