package logs

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
)

func GetContainerLogs() error {
	logsDir := "/logs_stor"
	files, err := ioutil.ReadDir(logsDir)
	if err != nil {
		log.Fatalf("Failed to read logs directory: %v", err)
		return err
	}
	for _, file := range files {
		if !file.IsDir() && filepath.Ext(file.Name()) == ".log" {
			logFilePath := filepath.Join(logsDir, file.Name())
			content, err := ioutil.ReadFile(logFilePath)
			if err != nil {
				log.Printf("Failed to read log file %s: %v", logFilePath, err)
				continue
			}
			fmt.Printf("Logs from %s:\n%s\n", file.Name(), content)
		}
	}
	return nil
}
