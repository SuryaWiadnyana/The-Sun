package main

import (
	"fmt"
)

type Buku struct {
	Judul   string
	Penulis string
	Halaman string
	Status  string
}

var daftarBuku []Buku

func tambahBuku(judul, penulis, halaman string) {
	buku := Buku{
		Judul:   judul,
		Penulis: penulis,
		Halaman: halaman,
		Status:  "tersedia",
	}
	daftarBuku = append(daftarBuku, buku)
	fmt.Println("Buku berhasil ditambahkan.")
	fmt.Println("=======================================")
}

func tampilkanBuku() {
	if len(daftarBuku) == 0 {
		fmt.Println("Tidak ada buku dalam daftar.")
		return
	}
	for i, buku := range daftarBuku {
		fmt.Printf("%d. Judul: %s, Penulis: %s, Halaman: %s, Status: %s\n", i+1, buku.Judul, buku.Penulis, buku.Halaman, buku.Status)
	}
}

func updateStatusBuku(index int, status string) {
	if index < 0 || index >= len(daftarBuku) {
		fmt.Println("Indeks buku tidak valid.")
		return
	}
	daftarBuku[index].Status = status
	fmt.Println("Status buku berhasil diperbarui.")
}

func hapusDataBuku(index int) {
	if index < 0 || index >= len(daftarBuku) {
		fmt.Println("Indeks buku tidak valid.")
		return
	}
	daftarBuku = append(daftarBuku[:index], daftarBuku[index+1:]...)
	fmt.Println("Buku berhasil dihapus.")
	fmt.Println("=======================================")
}

func main() {
	var pilihan int

	for {
		fmt.Println("\nManajemen Daftar Buku Perpustakaan")
		fmt.Println("1. Tambah Buku Baru")
		fmt.Println("2. Tampilkan Semua Buku")
		fmt.Println("3. Update Status Buku")
		fmt.Println("4. Hapus Buku")
		fmt.Println("5. Keluar")
		fmt.Print("Masukkan pilihan Anda: ")
		fmt.Scanln(&pilihan)

		switch pilihan {
		case 1:
			var judul, penulis, halaman string
			fmt.Print("Masukkan judul buku: ")
			fmt.Scanln(&judul)
			fmt.Print("Masukkan penulis buku: ")
			fmt.Scanln(&penulis)
			fmt.Print("Masukkan halaman buku: ")
			fmt.Scanln(&halaman)
			tambahBuku(judul, penulis, halaman)
		case 2:
			tampilkanBuku()
		case 3:
			var index int
			var status string
			fmt.Print("Masukkan nomor buku yang akan diupdate: ")
			fmt.Scanln(&index)
			fmt.Print("Masukkan status baru (tersedia/dipinjam): ")
			fmt.Scanln(&status)
			updateStatusBuku(index-1, status)
		case 4:
			tampilkanBuku()
			var index int
			fmt.Print("Masukkan nomor buku yang akan dihapus: ")
			fmt.Scanln(&index)
			hapusDataBuku(index - 1)
		case 5:
			fmt.Println("Terima kasih telah menggunakan program ini.")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}
