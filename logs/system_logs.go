package logs

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

type SysLogEntry struct {
	Timestamp string `json:"timestamp,omitempty"`
	Message   string `json:"message"`
}

func GetSystemdLogs() ([]SysLogEntry, error) {
	cmd := exec.Command("journalctl", "--no-pager", "--output=short-iso")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return nil, fmt.Errorf("error running journalctl: %v", err)
	}

	lines := strings.Split(out.String(), "\n")
	var logEntries []SysLogEntry
	for _, line := range lines {
		if line == "" {
			continue
		}
		logEntries = append(logEntries, parseLogLine(line))
	}
	return logEntries, nil
}
func parseLogLine(line string) SysLogEntry {
	parts := strings.SplitN(line, " ", 2)
	if len(parts) == 2 {
		return SysLogEntry{
			Timestamp: parts[0],
			Message:   parts[1],
		}
	}
	return SysLogEntry{Message: line}
}
