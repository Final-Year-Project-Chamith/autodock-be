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

// LogEntry represents a structured log entry for JSON output
type LogEntry struct {
	Timestamp string `json:"timestamp,omitempty"`
	Message   string `json:"message"`
}

func GetDockerContainerLogs(containerId string) error {
	options := container.LogsOptions{ShowStdout: true, ShowStderr: true, Timestamps: true,Follow: false}
	out, err := Client.ContainerLogs(context.Background(), containerId, options)
	if err != nil {
		return err
	}
	defer out.Close()

	// Buffer for reading logs
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

		// Split logs by line and parse each line
		lines := strings.Split(string(buf[:n]), "\n")
		for _, line := range lines {
			if line == "" {
				continue
			}
			logEntries = append(logEntries, parseLogLine(line))
		}
	}

	// Write the collected logs to a JSON file
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

// parseLogLine parses a single log line into a LogEntry
func parseLogLine(line string) LogEntry {
	// Split log line into timestamp and message
	parts := strings.SplitN(line, " ", 2)
	if len(parts) == 2 {
		return LogEntry{
			Timestamp: parts[0],
			Message:   parts[1],
		}
	}
	return LogEntry{Message: line}
}
