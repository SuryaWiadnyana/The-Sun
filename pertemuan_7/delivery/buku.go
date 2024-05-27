package delivery

import (
	"context"
	"fmt"
	"net/http"
	"pertemuan_7/domain"

	"github.com/asaskevich/govalidator"
	"github.com/gofiber/fiber/v2"
)

type HttpDeliveryBuku struct {
	HTTP domain.ServiceBuku
}

func NewHttpDeliveryBuku(app fiber.Router, HTTP domain.ServiceBuku) {
	handler := HttpDeliveryBuku{
		HTTP: HTTP,
	}

	group := app.Group("buku")
	group.Get("/All", handler.GetAll)
	group.Get("/by-id/:id", handler.GetBukuByID)
	group.Post("/create", handler.CreateBuku)
	group.Delete("/Delete/:id", handler.DeleteBuku)
	group.Put("/Update/:id", handler.UpdateBuku)
}

func (d *HttpDeliveryBuku) GetAll(c *fiber.Ctx) error {
	val, err := d.HTTP.GetAll(context.Background())
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "Gagal Untuk Mendapatkan Buku",
		})
	}
	fmt.Println(val)

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"data":  val,
		"error": err,
	})
}

func (b *HttpDeliveryBuku) CreateBuku(c *fiber.Ctx) error {

	req := new(domain.Buku)
	if err := c.BodyParser(req); err != nil {
		return err
	}

	ValRes, error := govalidator.ValidateStruct(req)
	if !ValRes {
		return error
	}

	bdy := domain.Buku{
		Judul:   req.Judul,
		Penulis: req.Penulis,
		Halaman: req.Halaman,
		Status:  req.Status,
	}

	data, err := b.HTTP.CreateBuku(context.Background(), &bdy)
	if err != nil {
		return err
	}

	return c.Status(http.StatusUnprocessableEntity).JSON(fiber.Map{
		"message": "Success",
		"data":    data,
	})
}

func (b *HttpDeliveryBuku) DeleteBuku(c *fiber.Ctx) error {
	id := c.Params("id")

	err := b.HTTP.DeleteBuku(context.Background(), id)
	if err != nil {
		errSign := err.Error()[0:1]
		if errSign == "^" {
			fmt.Println(errSign)
		}
		return c.Status(http.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": "Gagal Untuk Menghapus Buku",
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Buku Berhasil Dihapus",
	})

}

func (b *HttpDeliveryBuku) UpdateBuku(c *fiber.Ctx) error {

	id := c.Params("id")

	err := b.HTTP.UpdateBuku(context.Background(), &domain.Buku{})
	if err != nil {
		errSign := err.Error()[0:1]
		if errSign == "^" {
			fmt.Println(errSign)
		}
		return c.Status(http.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": "Gagal Untuk Menghapus Buku",
		})
	}

	body := new(domain.Buku)
	if err := c.BodyParser(body); err != nil {
		return err
	}

	ValRes, error := govalidator.ValidateStruct(body)
	if !ValRes {
		fmt.Println(c, "Failed to parses body", error)
	}

	bdy := domain.Buku{
		ID:      id,
		Judul:   body.Judul,
		Penulis: body.Penulis,
		Halaman: body.Halaman,
		Status:  body.Status,
	}

	err = b.HTTP.UpdateBuku(context.Background(), &bdy)
	if err != nil {
		return err
	}

	return c.Status(http.StatusUnprocessableEntity).JSON(fiber.Map{
		"message": "success",
	})
}

func (b *HttpDeliveryBuku) GetBukuByID(c *fiber.Ctx) error {
	id := c.Params("id")

	buku, err := b.HTTP.GetBukuByID(context.Background(), id)
	if err != nil {
		errSign := err.Error()[0:1]
		if errSign == "^" {
			fmt.Println(errSign)
		}
		return c.Status(http.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": "Gagal Untuk Menghapus Buku",
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "success",
		"data":    buku,
	})
}
