package api

import (
	"autodock-be/functions"
	"encoding/json"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func WebHookHandlerApi(c *fiber.Ctx) error {
	if c.Method() != fiber.MethodPost {
		return c.Status(fiber.StatusMethodNotAllowed).SendString("Invalid request method")
	}
	body := c.Body()

	signature := c.Get("X-Hub-Signature")
	if !functions.ValidateSignature(body, signature, "Asd@1234") {
		return c.Status(fiber.StatusForbidden).SendString("Invalid signature")
	}
	var payload map[string]interface{}
	if err := json.Unmarshal(body, &payload); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Failed to parse JSON")
	}
	fmt.Printf("Received webhook event: %v\n", payload)

	
	return c.SendString("Webhook received successfully")
}
