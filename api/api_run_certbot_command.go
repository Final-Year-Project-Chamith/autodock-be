package api

import (
	"autodock-be/functions"

	"github.com/gofiber/fiber/v2"
)

func RunCertbotCmdApi(c *fiber.Ctx) error {
	domain := c.Query("domain")

	if err := functions.RunCertbotManualDNS(domain); err != nil {
		return c.Status(fiber.StatusOK).JSON(map[string]string{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(map[string]string{"status": "certbot command successfull"})
}
