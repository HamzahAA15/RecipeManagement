package main

import (
	"log"
	"recipeApps/internals/httpservice"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	httpservice.InitRecipeRoute(app)
	httpservice.InitUserRoute(app)
	err := app.Listen(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
