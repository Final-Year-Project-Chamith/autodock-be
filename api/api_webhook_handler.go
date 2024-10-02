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

	eventType := c.Get("X-GitHub-Event")
	fmt.Printf("Received GitHub event: %s\n", eventType)

	if eventType == "pull_request" {
		action, ok := payload["action"].(string)
		if !ok {
			return c.Status(fiber.StatusBadRequest).SendString("Invalid pull request event structure")
		}

		if action == "closed" {
			pullRequest, ok := payload["pull_request"].(map[string]interface{})
			if !ok {
				return c.Status(fiber.StatusBadRequest).SendString("Invalid pull request data")
			}

			merged, ok := pullRequest["merged"].(bool)
			if ok && merged {

				baseBranch := pullRequest["base"].(map[string]interface{})["ref"].(string)
				if baseBranch == "main" {

					fmt.Printf("Pull request merged into main branch: %v\n", pullRequest)
				}
			}
		}
	}
	if eventType == "workflow_run" {
		workflowRun, ok := payload["workflow_run"].(map[string]interface{})
		if !ok {
			return c.Status(fiber.StatusBadRequest).SendString("Invalid workflow run data")
		}
		fmt.Println(workflowRun)

		if conclusion, ok := workflowRun["conclusion"].(string); ok {
			if conclusion == "success" {
				fmt.Printf("Workflow run successful: %v\n", workflowRun)
			} else {
				fmt.Printf("Workflow run failed or other state: %v\n", workflowRun)
			}
		} else {
			fmt.Println("Conclusion not found or is nil in workflow run.")
		}
	}


	return c.SendString("Webhook received successfully")
}
