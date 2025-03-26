package api

import (
	"autodock-be/functions"

	"github.com/gofiber/fiber/v2"
)

func RunCertbotCmdApi(c *fiber.Ctx) error {
	domain := c.Query("domain")
	//email := c.Query("email")

	result, err := functions.GenerateSSLCert(domain)
	if err != nil{
		return c.Status(fiber.StatusOK).JSON(map[string]string{"error":err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(result)
}
