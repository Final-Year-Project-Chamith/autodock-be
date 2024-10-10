package api

import (
	"autodock-be/docker"
	"autodock-be/dto"
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

	eventType := c.Get("X-GitHub-Event")
	fmt.Printf("Received GitHub event: %s\n", eventType)

	if eventType == "pull_request" {
		var pullRequestEvent dto.PullRequestEvent
		if err := json.Unmarshal(body, &pullRequestEvent); err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("Failed to parse pull request JSON")
		}

		if pullRequestEvent.Action == "closed" && pullRequestEvent.PullRequest.Merged {
			baseBranch := pullRequestEvent.PullRequest.Base.Ref
			if baseBranch == "main" {
				fmt.Printf("Pull request merged into main branch: %+v\n", pullRequestEvent.PullRequest)
			}
		}
	}

	if eventType == "workflow_run" {
		var workflowRunEvent dto.WorkflowRunEvent
		if err := json.Unmarshal(body, &workflowRunEvent); err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("Failed to parse workflow run JSON")
		}
		fmt.Println(workflowRunEvent.WorkflowRun.Conclusion)
		if workflowRunEvent.WorkflowRun.Conclusion == "success" {
			if err := docker.RunDockerComposeDeatched("D:\\Chamith\\Repos\\msg-app\\docker-compose.yml"); err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(map[string]string{"error": err.Error()})
			}
			fmt.Printf("Workflow run successful: %+v\n", workflowRunEvent.WorkflowRun)
		} else {
			fmt.Printf("Workflow run failed or other state: %+v\n", workflowRunEvent.WorkflowRun)
		}
	}

	return c.SendString("Webhook received successfully")
}
