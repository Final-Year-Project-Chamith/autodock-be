package main

import (
	"autodock-be/apiHandlers"
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

	// Connect To Database
	// dbConfig.ConnectToMongoDB()

	// //Remove Pre-Generated Outs
	// dbConfig.RemoveGeneratedOuts()

	// Define the API routes
	apiHandlers.Router(app)

	// Start the server
	log.Fatal(app.Listen(":8888"))
}
