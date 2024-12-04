package logs

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type SysLogEntry struct {
	Timestamp string `json:"timestamp,omitempty"`
	Message   string `json:"message"`
}

func GetSystemdLogs() ([]SysLogEntry, error) {

	logFile := "host_logs/messages"
	file, err := os.Open(logFile)
	if err != nil {
		return nil, fmt.Errorf("error opening log file: %v", err)
	}
	defer file.Close()

	var logEntries []SysLogEntry
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		logEntries = append(logEntries, parseLogLine(line))
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading log file: %v", err)
	}

	return logEntries, nil
}

func parseLogLine(line string) SysLogEntry {

	parts := strings.SplitN(line, " ", 3)
	if len(parts) == 3 {
		return SysLogEntry{
			Timestamp: fmt.Sprintf("%s %s", parts[0], parts[1]),
			Message:   parts[2],
		}
	}
	return SysLogEntry{Message: line}
}
