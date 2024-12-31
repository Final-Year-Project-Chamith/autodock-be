package api

import (
	"autodock-be/functions"
	"autodock-be/git"
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

	eventType := c.Get("X-GitHub-Event")
	fmt.Printf("Received GitHub event: %s\n", eventType)

	if err := git.HandleEventType(body, eventType); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	return c.SendString("Webhook received successfully")
}
