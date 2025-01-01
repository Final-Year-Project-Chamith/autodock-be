package dto

type SysLogEntry struct {
	Timestamp string `json:"timestamp,omitempty"`
	Message   string `json:"message"`
}

type ContainerLogs []struct {
	Log    string `json:"log,omitempty"`
	Stream string `json:"stream"`
	Time   string `json:"time"`
}
