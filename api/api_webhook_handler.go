package api

import (
	"autodock-be/dto"
	"autodock-be/functions"
	"autodock-be/git"

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

	eventType := c.Get("X-GitHub-Event")
	fmt.Printf("Received GitHub event: %s\n", eventType)

	switch eventType {
	case "pull_request":
		var pullRequestEvent dto.PullRequestEvent
		if err := json.Unmarshal(body, &pullRequestEvent); err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("Failed to parse pull request JSON")
		}
		if err := git.HandlePullRequest(pullRequestEvent); err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}
	case "workflow_run":
		var workflowRunEvent dto.WorkflowRunEvent
		if err := json.Unmarshal(body, &workflowRunEvent); err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("Failed to parse workflow run JSON")
		}
		if err := git.HandleWorkflowRun(workflowRunEvent); err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}
	default:
		return c.SendString("Unhandled event type")
	}

	return c.SendString("Webhook received successfully")
}
