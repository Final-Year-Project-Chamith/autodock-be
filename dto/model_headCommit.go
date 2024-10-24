package dto

type HeadCommit struct {
	Author    Author `json:"author"`
	Committer Author `json:"committer"`
	ID        string `json:"id"`
	Message   string `json:"message"`
	Timestamp string `json:"timestamp"`
	TreeID    string `json:"tree_id"`
}
