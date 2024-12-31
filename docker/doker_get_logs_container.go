package docker

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/docker/docker/api/types/container"
)

type LogEntry struct {
	Timestamp  string `json:"timestamp,omitempty"`
	LogLevel   string `json:"log_level,omitempty"`
	ContainerID string `json:"container_id,omitempty"`
	Message    string `json:"message"`
}


func GetDockerContainerLogs(containerId string) error {
	options := container.LogsOptions{ShowStdout: true, ShowStderr: true, Timestamps: true,Follow: false}
	out, err := Client.ContainerLogs(context.Background(), containerId, options)
	if err != nil {
		return err
	}
	defer out.Close()
	var logEntries []LogEntry
	buf := make([]byte, 4096)

	for {
		n, err := out.Read(buf)
		if err != nil && err != io.EOF {
			return fmt.Errorf("error reading logs: %v", err)
		}
		if n == 0 {
			break
		}
		lines := strings.Split(string(buf[:n]), "\n")
		for _, line := range lines {
			if line == "" {
				continue 
			}
			logEntries = append(logEntries, parseLogLine(line))
		}
	}
	file, err := os.Create("container_logs.json")
	if err != nil {
		return fmt.Errorf("error creating JSON file: %v", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(logEntries); err != nil {
		return fmt.Errorf("error encoding logs to JSON: %v", err)
	}

	fmt.Println("Logs successfully saved to container_logs.json")
	return nil
}


func parseLogLine(line string) LogEntry {
	fmt.Println("Raw Log Line:", line)
	parts := strings.Fields(line)
	logEntry := LogEntry{}

	if len(parts) > 0 {
		logEntry.Timestamp = parts[0] 
	}
	if len(parts) > 1 {
		logEntry.LogLevel = parts[1] 
	}
	if len(parts) > 2 && strings.HasPrefix(parts[2], "Container") {
		logEntry.ContainerID = strings.TrimPrefix(parts[2], "Container")
	}
	if len(parts) > 3 {
		logEntry.Message = strings.Join(parts[3:], " ")
	}

	return logEntry
}
