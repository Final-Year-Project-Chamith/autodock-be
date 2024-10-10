package dto

type PullRequest struct {
	ID     int    `json:"id"`
	Number int    `json:"number"`
	State  string `json:"state"`
	Merged bool   `json:"merged"`
	Base   Branch `json:"base"`
}

type Branch struct {
	Ref string `json:"ref"`
}

type PullRequestEvent struct {
	Action      string      `json:"action"`
	PullRequest PullRequest `json:"pull_request"`
}
