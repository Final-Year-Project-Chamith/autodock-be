package api

import (
	"autodock-be/dto"
	"autodock-be/functions"

	"github.com/gofiber/fiber/v2"
)

func GenerateNginxFile(c *fiber.Ctx)error{
	inoutObj := dto.NginxConf{}
	if err := c.BodyParser(&inoutObj); err != nil{
		return c.Status(fiber.StatusBadRequest).JSON(map[string]string{"error":err.Error()})
	}
	err := functions.GenerateNginxFile(inoutObj)
	if err != nil{
		return c.Status(fiber.StatusInternalServerError).JSON(map[string]string{"error":err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(map[string]string{"error":"nginx file successfully generated"})
}