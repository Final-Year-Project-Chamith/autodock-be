package api

import (
	"autodock-be/functions"
	"fmt"
	"time"

	"github.com/gofiber/websocket/v2"
)

func GetSystemMetrices(c *websocket.Conn) {
	for {
		metrics, err := functions.GetUsageMetrices()
		if err != nil {
			fmt.Println("Error fetching metrics:", err)
			break
		}
		if err := c.WriteJSON(metrics); err != nil {
			fmt.Println("Error sending metrics:", err)
			break
		}
		time.Sleep(1 * time.Second)
	}
}
