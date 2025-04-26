package api

import (
	"autodock-be/docker"

	"github.com/gofiber/fiber/v2"
)

func GetContainerLogsApi(c *fiber.Ctx) error {
	containerId := c.Query("containerId")
	logs, err := docker.GetDockerContainerLogs(containerId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(map[string]string{"error": err.Error()})
	}
	return c.JSON(logs)
}
