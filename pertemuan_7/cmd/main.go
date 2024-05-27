package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"pertemuan_7/config"
	httpapi "pertemuan_7/delivery"
	"pertemuan_7/domain"
	"pertemuan_7/library/db"
	"pertemuan_7/repository"
	"pertemuan_7/services"
	"runtime/debug"
	"syscall"
	"time"

	"github.com/bytedance/sonic"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

var ServiceBuku domain.ServiceBuku

func bootDatabase() {

	godotenv.Load("../.env")
	db.InitMongoDB()
}

func bootServices() {
	ucTimeout := 60 * time.Second

	RepoBuku := repository.NewMongoRepoBuku(db.Mongodb)
	ServiceBuku = services.NewServiceBuku(RepoBuku, ucTimeout)
}
func bootstrapFiber() *fiber.App {
	app := fiber.New(
		fiber.Config{
			DisableStartupMessage: false,
			JSONEncoder:           sonic.Marshal,
			JSONDecoder:           sonic.Unmarshal,
			Prefork:               false,
			ServerHeader:          "BTW Edutech",
			AppName:               config.GetAppName(),
			ReadTimeout:           time.Second * 60,
			CaseSensitive:         true,
			ErrorHandler: func(c *fiber.Ctx, err error) error {
				// Status code defaults to 500
				code := fiber.StatusInternalServerError

				// Retrieve the custom status code if it's an fiber.*Error
				if e, ok := err.(*fiber.Error); ok {
					code = e.Code
				}

				if code == 404 {
					fmt.Println("Error 404")
				}

				// Return from handler
				return nil
			},
		},
	)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).SendString("OK")
	})

	app.Use(

		func(c *fiber.Ctx) error {
			defer func() {
				if r := recover(); r != nil {
					// TODO:
					id := time.Now().Unix()

					fmt.Printf("Server Panic Occured [CrashID: %d]\n[PanicMessage]:%s\n[PanicStack]:%s\n", id, r, string(debug.Stack()))
					fmt.Printf("Server Panic Occured [CrashID: %d]\n[Body]:%s\n[ResponseBody]:%s\n[Header]:%s", id, string(c.Body()), string(c.Response().Body()), &c.Request().Header)

					c.Status(fiber.StatusInternalServerError).SendString("Server Error")
				}
			}()
			return c.Next()
		},

		// func(c *fiber.Ctx) error {
		// 	c.Locals("wg", wg)

		// 	return c.Next()
		// },
	)
	return app
}

func main() {
	bootDatabase()
	bootServices()
	startHttp := flag.Bool("start-http", true, "Start Http Server")
	flag.Parse()
	if *startHttp {
		initHttp()
	}
}

func initHttp() {
	app := bootstrapFiber()
	rootRouter := app.Group("/")

	httpapi.NewHttpDeliveryBuku(rootRouter, ServiceBuku)

	fmt.Println("HTTP App Start and Running")

	go func() {
		if err := app.Listen(":8000"); err != nil {
			fmt.Println("Error starting location HTTP server: ", err)
		}
	}()
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)
	<-interrupt

	fmt.Println("HTTP Shutting down")
	if err := app.Shutdown(); err != nil {
		fmt.Println(err)
	}
}
