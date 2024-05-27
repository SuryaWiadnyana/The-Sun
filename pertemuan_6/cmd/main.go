package main

import (
	"fmt"
	"pertemuan_6/Model"

	"github.com/gofiber/fiber/v2"
)

var bookList []Model.Books

func init() {
	bookList = append(bookList, Model.Books{ID: "01", Judul: "Mentari", Penulis: "Surya", Halaman: 85, Status: "Tersedia"})
	bookList = append(bookList, Model.Books{ID: "02", Judul: "Bulan", Penulis: "Surya", Halaman: 100, Status: "Tersedia"})
	bookList = append(bookList, Model.Books{ID: "03", Judul: "Surya", Penulis: "Surya", Halaman: 50, Status: "Tersedia"})
}

func getAll(c *fiber.Ctx) error {
	return c.JSON(bookList)
}

func getBooksByID(c *fiber.Ctx) error {
	id := c.Params("id")

	for _, books := range bookList {
		if books.ID == id {
			return c.JSON(books)
		}
	}

	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
		"message": "Book not found",
	})
}

func createBooks(c *fiber.Ctx) error {
	var books Model.Books
	if err := c.BodyParser(&books); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	bookList = append(bookList, books)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Success Insert Data",
		"data":    books,
	})
}

func updateBooks(c *fiber.Ctx) error {
	id := c.Params("id")

	var req Model.Books
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	for i, books := range bookList {
		if books.ID == id {
			bookList[i] = Model.Books{
				ID:      req.ID,
				Judul:   req.Judul,
				Penulis: req.Penulis,
				Halaman: req.Halaman,
				Status:  req.Status,
			}
			return c.JSON(fiber.Map{
				"message": "Book updated successfully",
				"data":    bookList[i],
			})
		}
	}

	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
		"message": "Book not found",
	})
}

func deleteBooks(c *fiber.Ctx) error {
	id := c.Params("id")

	for i, books := range bookList {
		if books.ID == id {
			bookList = append(bookList[:i], bookList[i+1:]...)
			return c.JSON(fiber.Map{
				"message": "Book deleted successfully",
			})
		}
	}

	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
		"message": "Book not found",
	})
}

func SearchBooks(c *fiber.Ctx) error {
	query := c.Query("judul")

	var results []Model.Books
	for _, books := range bookList {
		if books.Judul == query {
			results = append(results, books)
		}
	}

	if len(results) == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "No books found",
		})
	}

	return c.JSON(results)
}

func main() {
	app := fiber.New()

	BooksRoute := app.Group("/Books")
	BooksRoute.Get("/All", getAll)
	BooksRoute.Get("/by-id/:id", getBooksByID)
	BooksRoute.Post("/create", createBooks)
	BooksRoute.Put("/update/:id", updateBooks)
	BooksRoute.Delete("/delete/:id", deleteBooks)
	BooksRoute.Get("/search", SearchBooks)

	err := app.Listen(":8000")
	if err != nil {
		fmt.Println("ERROR: ", err)
	}
}
