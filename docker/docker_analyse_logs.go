package docker

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/docker/docker/api/types/container"
)

type AnalysisResponse struct {
	Report       string `json:"report"`
	Model        string `json:"model"`
	AnalysisType string `json:"analysis_type"`
	Timestamp    string `json:"timestamp"`
}

// Get logs and analyze them with the LLM
func AnalyzeContainerLogs(containerId string) (*AnalysisResponse, error) {
	logs, err := getLogs(containerId)
	if err != nil {
		return nil, err
	}
	fmt.Println(logs)
	payload := map[string]interface{}{
		"logs": logs,
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}
	fmt.Println(jsonData)
	resp, err := http.Post("https://bc5e-2402-d000-812c-60e-617b-f900-dfe-9906.ngrok-free.app/analyze", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("failed to contact analyzer API: %v", err)
	}
	defer resp.Body.Close()

	var analysis AnalysisResponse
	err = json.NewDecoder(resp.Body).Decode(&analysis)
	if err != nil {
		return nil, fmt.Errorf("failed to decode analysis response: %v", err)
	}

	return &analysis, nil
}


func getLogs(containerId string) ([]LogEntry, error) {
	options := container.LogsOptions{
		ShowStdout: true,
		ShowStderr: true,
		Timestamps: true,
		Follow:     false,
	}

	out, err := Client.ContainerLogs(context.Background(), containerId, options)
	if err != nil {
		return nil, err
	}
	defer out.Close()

	var logEntries []LogEntry
	buf := make([]byte, 4096)

	for {
		n, err := out.Read(buf)
		if err != nil && err != io.EOF {
			return nil, fmt.Errorf("error reading logs: %v", err)
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

	// ✅ Limit to the first 4 log lines
	if len(logEntries) > 4 {
		logEntries = logEntries[:4]
	}

	return logEntries, nil
}
