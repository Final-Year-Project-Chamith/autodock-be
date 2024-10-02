package api

import (
	"autodock-be/docker"

	"github.com/gofiber/fiber/v2"
)

func DeployHandlerApi(c *fiber.Ctx)error{
	if err := docker.RunDockerComposeDeatched("D:\\Chamith\\Repos\\msg-app\\docker-compose.yml"); err != nil{
		return c.Status(fiber.StatusInternalServerError).JSON(map[string]string{"error":err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(map[string]string{"error":"successfully deployed"})
}