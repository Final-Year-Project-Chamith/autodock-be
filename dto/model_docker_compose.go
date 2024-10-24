package dto

type DockerCompose struct {
	Repo string `json:"repo"`
	Services []Service `json:"services"`
}

type Service struct {
	Name      string            `json:"name"`
	Image     string            `json:"image"`
	Container string            `json:"container"`
	Ports     []string          `json:"ports"`
	Volumes   []string          `json:"volumes"`
	EnvVars   map[string]string `json:"envVars"`
}
