package api

import (
	"autodock-be/docker"
	"autodock-be/dto"
	"autodock-be/functions"
	"errors"

	"github.com/gofiber/fiber/v2"
)

func GenerateDockerComposeFileApi(c *fiber.Ctx) error {
	inputObj := dto.DockerCompose{}
	if err := c.BodyParser(&inputObj); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(map[string]string{"error": err.Error()})
	}
	if err := functions.GenerateDockerComposeFile(inputObj, inputObj.Repo); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(map[string]string{"error": err.Error()})
	}
	if err := docker.RunDockerComposeDeatched("docker-compose/" + inputObj.Repo + "/docker-compose.yml"); err != nil {
		return errors.New("failed to run docker compose: " + err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(map[string]string{"status": "Success"})
}
