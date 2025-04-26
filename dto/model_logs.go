package dto

type SysLogEntry struct {
	Timestamp string `json:"timestamp,omitempty"`
	Message   string `json:"message"`
}

type ContainerLog struct {
	ContainerID string `json:"container_id"`
	LogDetails  struct {
		Log    string `json:"log,omitempty"`
		Stream string `json:"stream"`
		Time   string `json:"time"`
	} `json:"log"`
}
type LogEntry struct {
	Timestamp   string `json:"timestamp,omitempty"`
	LogLevel    string `json:"log_level,omitempty"`
	ContainerID string `json:"container_id,omitempty"`
	Message     string `json:"message"`
}
