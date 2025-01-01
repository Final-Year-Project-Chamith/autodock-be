package logs

import (
	"autodock-be/dto"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

func GetContainerLogs() (interface{}, error) {
	logsFile := "/app/logs_stor/all_containers_logs.json"

	// Debug: Check the directory structure
	log.Println("Checking logs_stor directory:")
	files, err := ioutil.ReadDir("/app/logs_stor/")
	if err != nil {
		log.Printf("Failed to read directory: %v", err)
		return nil, fmt.Errorf("failed to read logs directory: %w", err)
	}
	for _, f := range files {
		log.Printf("Found file: %s", f.Name())
	}

	// Read the JSON file
	content, err := ioutil.ReadFile(logsFile)
	if err != nil {
		log.Printf("Failed to read log file: %v", err)
		return nil, fmt.Errorf("log file not found: %w", err)
	}

	// Parse the JSON content
	var logs []dto.ContainerLogs
	if err := json.Unmarshal(content, &logs); err != nil {
		log.Printf("Failed to parse JSON: %v", err)
		return nil, fmt.Errorf("invalid JSON format: %w", err)
	}

	return logs, nil
}
