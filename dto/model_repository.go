package dto

type Repository struct {
	ArchiveURL    string `json:"archive_url"`
	AssigneesURL  string `json:"assignees_url"`
	BlobsURL      string `json:"blobs_url"`
	BranchesURL   string `json:"branches_url"`
	CollaboratorsURL string `json:"collaborators_url"`
	CommentsURL   string `json:"comments_url"`
	CommitsURL    string `json:"commits_url"`
	CompareURL    string `json:"compare_url"`
	ContentsURL   string `json:"contents_url"`
	ContributorsURL string `json:"contributors_url"`
	DeploymentsURL string `json:"deployments_url"`
	Description   string `json:"description"`
	DownloadsURL  string `json:"downloads_url"`
	EventsURL     string `json:"events_url"`
	Fork          bool   `json:"fork"`
	ForksURL      string `json:"forks_url"`
	FullName      string `json:"full_name"`
	GitCommitsURL string `json:"git_commits_url"`
	GitRefsURL    string `json:"git_refs_url"`
	GitTagsURL    string `json:"git_tags_url"`
	HTMLURL       string `json:"html_url"`
	ID            int64  `json:"id"`
}