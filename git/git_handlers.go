package git

import (
	"autodock-be/docker"
	"autodock-be/dto"
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
		if err := docker.RunDockerComposeDeatched("D:\\Chamith\\Repos\\msg-app\\docker-compose.yml"); err != nil {
			return errors.New("failed to run docker compose: " + err.Error())
		}
		fmt.Printf("Workflow run successful: %+v\n", event.WorkflowRun)
	} else {
		fmt.Printf("Workflow run failed or other state: %+v\n", event.WorkflowRun)
	}
	return nil
}
