package api

import (
	"autodock-be/docker"

	"github.com/gofiber/fiber/v2"
)

func GetAllDockerContainers(c *fiber.Ctx)error{
	containers, err := docker.ListAllContainers()
	if err != nil{
		return c.Status(fiber.StatusInternalServerError).JSON(map[string]string{"error":err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(containers)
}