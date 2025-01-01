package api

import (
	"github.com/gofiber/fiber/v2"
)

func GetContainerLogsApi(c *fiber.Ctx) error {
	//containerId := c.Query("containerId")
	// _, err := logs.GetContainerLogs()
	// if err != nil {
	// 	return c.Status(fiber.StatusInternalServerError).JSON(map[string]string{"error": err.Error()})
	// }
	return c.Status(fiber.StatusOK).JSON("/logs_stor/all_containers_logs.json")
}
 