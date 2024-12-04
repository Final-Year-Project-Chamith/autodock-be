package api

import (
	"autodock-be/logs"
	"autodock-be/redis"
	"encoding/json"
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
)

func GetSystemLogs(c *fiber.Ctx) error {
	logs, err := logs.GetSystemdLogs()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(map[string]string{"error": err.Error()})
	}
	if err := redis.SaveLogsToRedis(redis.RedClient, logs); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(map[string]string{"error": err.Error()})
	}
	outFile, err := os.Create("system_logs.json")
	if err != nil {
		fmt.Printf("Error creating output file: %v\n", err)
		return c.Status(fiber.StatusInternalServerError).JSON(map[string]string{"error": err.Error()})
	}
	defer outFile.Close()

	encoder := json.NewEncoder(outFile)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(logs); err != nil {
		fmt.Printf("Error encoding logs to JSON: %v\n", err)
	} else {
		fmt.Println("System logs saved to system_logs.json")
	}
	return c.Status(fiber.StatusOK).JSON(map[string]string{"status": "system logs generated successfully"})
}
