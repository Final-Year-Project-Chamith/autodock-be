package git

import (
	"autodock-be/docker"
	"autodock-be/dto"
	"encoding/json"
	"errors"
	"fmt"
)

func HandlePullRequest(event dto.PullRequestEvent) error {
	if event.Action == "closed" && event.PullRequest.Merged {
		baseBranch := event.PullRequest.Base.Ref
		if baseBranch == "main" {
			fmt.Printf("Pull request merged into main branch: %+v\n", event.PullRequest)
		}
	}
	return nil
}

func HandleWorkflowRun(event dto.WorkflowRunEvent) error {
	if event.WorkflowRun.Conclusion == "success" {
		fmt.Println(event.WorkflowRun.Repository.FullName)
		fmt.Println("docker-compose/"+event.WorkflowRun.Repository.FullName+"/docker-compose.yml")
		if err := docker.RunDockerComposeDeatched("docker-compose/"+event.WorkflowRun.Repository.FullName+"/docker-compose.yml"); err != nil {
			return errors.New("failed to run docker compose: " + err.Error())
		}
		fmt.Printf("Workflow run successful: %+v\n", event.WorkflowRun)
	} else {
		fmt.Printf("Workflow run failed or other state: %+v\n", event.WorkflowRun)
	}
	return nil
}

func HandleEventType(body []byte, eventType string) error {
	fmt.Println(eventType)
	switch eventType {
	case "pull_request":
		var pullRequestEvent dto.PullRequestEvent
		if err := json.Unmarshal(body, &pullRequestEvent); err != nil {
			return err
		}
		if err := HandlePullRequest(pullRequestEvent); err != nil {
			return err
		}
	case "workflow_run":
		var workflowRunEvent dto.WorkflowRunEvent
		if err := json.Unmarshal(body, &workflowRunEvent); err != nil {
			return err
		}
		if err := HandleWorkflowRun(workflowRunEvent); err != nil {
			return err
		}
	case "ping":
		fmt.Println("godak ping")
	default:
		return fmt.Errorf("unhandled event type")
	}

	return nil

}
