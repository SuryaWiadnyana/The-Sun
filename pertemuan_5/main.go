package main

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Book struct {
	ID      int    `gorm:"primaryKey"`
	Judul   string `gorm:"type:varchar(100)"`
	Penulis string `gorm:"type:varchar(100)"`
	Halaman int
	Status  string `gorm:"type:varchar(20)"`
}

var DB *gorm.DB
var err error

func initDB() {
	dsn := "root:@tcp(127.0.0.1:3307)/library?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Gagal koneksi ke database: ", err)
	}

	DB.AutoMigrate(&Book{})
}

func createBook(book Book) {
	result := DB.Create(&book)
	if result.Error != nil {
		log.Println("Gagal menambahkan buku:", result.Error)
	} else {
		log.Println("Buku berhasil ditambahkan")
	}
}

func getAllBooks() []Book {
	var books []Book
	result := DB.Find(&books)
	if result.Error != nil {
		log.Println("Gagal mengambil data buku:", result.Error)
	}
	return books
}

func updateBookStatus(id int, status string) {
	var book Book
	result := DB.First(&book, id)
	if result.Error != nil {
		log.Println("Buku tidak ditemukan:", result.Error)
		return
	}
	book.Status = status
	DB.Save(&book)
	log.Println("Status buku berhasil diperbarui")
}

func deleteBook(id int) {
	var book Book
	result := DB.Delete(&book, id)
	if result.Error != nil {
		log.Println("Gagal menghapus buku:", result.Error)
	} else {
		log.Println("Buku berhasil dihapus")
	}
}

func clearScreen() {
	fmt.Print("\033[H\033[2J")
}

func main() {
	initDB()

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

		clearScreen()

		switch pilihan {
		case 1:
			var judul, penulis string
			var halaman int
			fmt.Print("Masukkan judul buku: ")
			fmt.Scanln(&judul)
			fmt.Print("Masukkan penulis buku: ")
			fmt.Scanln(&penulis)
			fmt.Print("Masukkan jumlah halaman: ")
			fmt.Scanln(&halaman)
			createBook(Book{Judul: judul, Penulis: penulis, Halaman: halaman, Status: "tersedia"})
		case 2:
			books := getAllBooks()
			for _, book := range books {
				fmt.Printf("ID: %d, Judul: %s, Penulis: %s, Halaman: %d, Status: %s\n",
					book.ID, book.Judul, book.Penulis, book.Halaman, book.Status)
			}
		case 3:
			var id int
			var status string
			fmt.Print("Masukkan ID buku yang akan diupdate: ")
			fmt.Scanln(&id)
			fmt.Print("Masukkan status baru (tersedia/dipinjam): ")
			fmt.Scanln(&status)
			updateBookStatus(id, status)
		case 4:
			var id int
			fmt.Print("Masukkan ID buku yang akan dihapus: ")
			fmt.Scanln(&id)
			deleteBook(id)
		case 5:
			fmt.Println("Terima kasih telah menggunakan program ini.")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}
