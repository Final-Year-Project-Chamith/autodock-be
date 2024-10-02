package main

import (
	"autodock-be/apiHandlers"
	"autodock-be/docker"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	fmt.Println("Starting application")

	app := fiber.New(fiber.Config{
		AppName:   "AutoDock-BE",
		BodyLimit: 4000 * 1024,
	})
	docker.ConnectDocker()

	

	
	apiHandlers.Router(app)

	
	log.Fatal(app.Listen(":8888"))
}
