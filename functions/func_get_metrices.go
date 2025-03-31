package functions

import (
	"autodock-be/dto"
	"fmt"
	"time"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/mem"
)

func GetUsageMetrices() (*dto.Metrics, error) {
	var metrics dto.Metrics
	cpuPercentage, err := cpu.Percent(time.Second, false)
	if err != nil {
		return nil, err
	}
	ramPercentage, err := mem.VirtualMemory()
	if err != nil {
		return nil, err
	}
	diskUsage, err := disk.Usage("/")
	if err != nil {
		return nil, err
	}
	fmt.Println("total ram", ramPercentage.Total)
	fmt.Println("total cpu", cpuPercentage[0])
	fmt.Println("total storage",diskUsage.Total)
	metrics.CPUUsage = cpuPercentage[0]*100
	metrics.MemoryUsage = ramPercentage.UsedPercent
	metrics.DiskUsage = diskUsage.UsedPercent

	return &metrics, nil
}
