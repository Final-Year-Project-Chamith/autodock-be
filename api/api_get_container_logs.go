package api

import (
	"autodock-be/logs"

	"github.com/gofiber/fiber/v2"
)

func GetContainerLogsApi(c *fiber.Ctx) error {
	//containerId := c.Query("containerId")
	if err := logs.GetContainerLogs(); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(map[string]string{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON("logs generated successfully")
}
