package redis

import (
	"autodock-be/dto"
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

func SaveLogsToRedis(client *redis.Client, logEntries []dto.SysLogEntry) error {
	for i, entry := range logEntries {
		key := fmt.Sprintf("syslog:%d", i) // Unique key for each log entry
		err := client.HSet(context.Background(), key, map[string]interface{}{
			"timestamp": entry.Timestamp,
			"message":   entry.Message,
		}).Err()
		if err != nil {
			return fmt.Errorf("error saving log to Redis: %v", err)
		}
	}
	return nil
}
