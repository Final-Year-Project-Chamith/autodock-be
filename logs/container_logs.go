package logs

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

func GetContainerLogs() (interface{}, error) {
	logsFile := "/logs_stor/all_containers_logs.json"

	// Read the JSON file
	content, err := ioutil.ReadFile(logsFile)
	if err != nil {
		log.Printf("Failed to read log file: %v", err)
		return nil, err
	}

	var logs interface{}
	if err := json.Unmarshal(content, &logs); err != nil {
		log.Printf("Failed to parse JSON: %v", err)
		return nil, err
	}
	return logs, nil
}
