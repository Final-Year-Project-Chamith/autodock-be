package dto

type WorkflowRun struct {
	Actor          Actor      `json:"actor"`
	ArtifactsURL   string     `json:"artifacts_url"`
	CancelURL      string     `json:"cancel_url"`
	CheckSuiteID   int64      `json:"check_suite_id"`
	CheckSuiteURL  string     `json:"check_suite_url"`
	Conclusion     string     `json:"conclusion"`
	CreatedAt      string     `json:"created_at"`
	DisplayTitle   string     `json:"display_title"`
	Event          string     `json:"event"`
	HeadBranch     string     `json:"head_branch"`
	HeadCommit     HeadCommit `json:"head_commit"`
	HeadRepository Repository `json:"head_repository"`
	HeadSHA        string     `json:"head_sha"`
	HTMLURL        string     `json:"html_url"`
	ID             int64      `json:"id"`
	JobsURL        string     `json:"jobs_url"`
	LogsURL        string     `json:"logs_url"`
	Name           string     `json:"name"`
	NodeID         string     `json:"node_id"`
	Path           string     `json:"path"`
	Repository     Repository `json:"repository"`
}

type WorkflowRunEvent struct {
	Action      string      `json:"action"`
	WorkflowRun WorkflowRun `json:"workflow_run"`
}
