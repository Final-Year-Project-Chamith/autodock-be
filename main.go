package main

import (
	"autodock-be/apiHandlers"
	"autodock-be/cmd"
	"autodock-be/docker"
	"autodock-be/functions"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/cobra"
)

func main() {
	if len(os.Args) > 1 {

		rootCmd := &cobra.Command{
			Use:   "autodock",
			Short: "A tool for managing Docker services",
		}
		functions.RemoveGeneratedOuts()
		rootCmd.AddCommand(cmd.GenerateDockerComposeCMD())

		if err := rootCmd.Execute(); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	} else {
		fmt.Println("Starting application")

		app := fiber.New(fiber.Config{
			AppName:   "AutoDock-BE",
			BodyLimit: 4000 * 1024,
		})
		functions.RemoveGeneratedOuts()
		docker.ConnectDocker()

		apiHandlers.Router(app)

		log.Fatal(app.Listen(":8888"))
	}
}
