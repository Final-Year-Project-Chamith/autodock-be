package api

import (
	"autodock-be/docker"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func AnalyzeContainerLogs(c *fiber.Ctx) error {
	containerId := c.Query("containerId")
	report, err := docker.AnalyzeContainerLogs(containerId)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("LLM Report:\n", report.Report)
	}
	return c.Status(fiber.StatusOK).JSON(report)
}
