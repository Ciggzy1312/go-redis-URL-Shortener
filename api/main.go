package main

import (
	"fmt"
	"log"

	"github.com/Ciggzy1312/url-shortener/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()

	if err != nil {
		fmt.Println(err)
	}

	app := fiber.New()

	app.Use(logger.New())

	app.Get(":/url", routes.ResolveURL)
	app.Post(":/api/v1", routes.ShortenURL)

	log.Fatal(app.Listen(":5000"))

}
