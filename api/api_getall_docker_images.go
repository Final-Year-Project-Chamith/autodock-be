package api

import (
	"autodock-be/docker"

	"github.com/gofiber/fiber/v2"
)

func GetAllDockerImages(c *fiber.Ctx)error{
	images, err := docker.ListAllDockerImages()
	if err != nil{
		return c.Status(fiber.StatusInternalServerError).JSON(map[string]string{"error":err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(images)
}