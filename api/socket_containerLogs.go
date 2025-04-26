package api

import (
	"autodock-be/docker"
	"autodock-be/dto"
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/docker/docker/api/types/container"
	"github.com/gofiber/websocket/v2"
)

func StreamContainerLogsWS(c *websocket.Conn) {
	containerId := c.Query("containerId") // expects /ws/logs/:id

	options := container.LogsOptions{
		ShowStdout: true,
		ShowStderr: true,
		Timestamps: true,
		Follow:     true,
		Tail:       "50",
	}

	logStream, err := docker.Client.ContainerLogs(context.Background(), containerId, options)
	if err != nil {
		c.WriteMessage(websocket.TextMessage, []byte("❌ Failed to get logs: "+err.Error()))
		return
	}
	defer logStream.Close()

	scanner := bufio.NewScanner(logStream)
	for scanner.Scan() {
		line := scanner.Text()
		entry := parseLogLine(line)
		entry.ContainerID = containerId

		jsonLog, _ := json.Marshal(entry)
		if err := c.WriteMessage(websocket.TextMessage, jsonLog); err != nil {
			break
		}
	}
}
func parseLogLine(line string) dto.LogEntry {
	fmt.Println("Raw Log Line:", line)
	parts := strings.Fields(line)
	logEntry := dto.LogEntry{}

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
