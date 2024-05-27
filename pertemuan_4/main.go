package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Buku struct {
	ID      int
	Judul   string
	Penulis string
	Halaman int
	Status  string
}

var db *sql.DB

func initDB() {
	var err error
	dsn := "root:@tcp(127.0.0.1:3307)/library"
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Koneksi ke database berhasil!")
}

func tambahBuku(judul, penulis string, halaman int) {
	_, err := db.Exec("INSERT INTO books (judul, penulis, halaman, status) VALUES (?, ?, ?, ?)", judul, penulis, halaman, "tersedia")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Buku berhasil ditambahkan.")
	time.Sleep(2 * time.Second)
	clearScreen()
}

func tampilkanBuku() {
	rows, err := db.Query("SELECT id, judul, penulis, halaman, status FROM books")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	fmt.Println("Daftar Buku:")
	for rows.Next() {
		var buku Buku
		err := rows.Scan(&buku.ID, &buku.Judul, &buku.Penulis, &buku.Halaman, &buku.Status)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%d. Judul: %s, Penulis: %s, Halaman: %d, Status: %s\n", buku.ID, buku.Judul, buku.Penulis, buku.Halaman, buku.Status)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Tekan Enter untuk melanjutkan...")
	fmt.Scanln()
	clearScreen()
}

func updateStatusBuku(id int, status string) {
	_, err := db.Exec("UPDATE books SET status = ? WHERE id = ?", status, id)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Status buku berhasil diperbarui.")
	time.Sleep(2 * time.Second)
	clearScreen()
}

func hapusDataBuku(id int) {
	_, err := db.Exec("DELETE FROM books WHERE id = ?", id)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Buku berhasil dihapus.")
	time.Sleep(2 * time.Second)
	clearScreen()
}

func clearScreen() {
	var clearCmd *exec.Cmd
	if runtime.GOOS == "windows" {
		clearCmd = exec.Command("cmd", "/c", "cls")
	} else {
		clearCmd = exec.Command("clear")
	}
	clearCmd.Stdout = os.Stdout
	clearCmd.Run()
}

func main() {
	initDB()
	defer db.Close()

	var pilihan int

	for {
		clearScreen()
		fmt.Println("Manajemen Daftar Buku Perpustakaan")
		fmt.Println("1. Tambah Buku Baru")
		fmt.Println("2. Tampilkan Semua Buku")
		fmt.Println("3. Update Status Buku")
		fmt.Println("4. Hapus Buku")
		fmt.Println("5. Keluar")
		fmt.Print("Masukkan pilihan Anda: ")
		fmt.Scanln(&pilihan)

		switch pilihan {
		case 1:
			var judul, penulis string
			var halaman int
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
			var id int
			var status string
			fmt.Print("Masukkan ID buku yang akan diupdate: ")
			fmt.Scanln(&id)
			fmt.Print("Masukkan status baru (tersedia/dipinjam): ")
			fmt.Scanln(&status)
			updateStatusBuku(id, status)
		case 4:
			var id int
			fmt.Print("Masukkan ID buku yang akan dihapus: ")
			fmt.Scanln(&id)
			hapusDataBuku(id)
		case 5:
			fmt.Println("Terima kasih telah menggunakan program ini.")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}
