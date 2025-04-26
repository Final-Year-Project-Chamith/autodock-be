package functions

import (
	"autodock-be/dto"
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
	metrics.CPUUsage = cpuPercentage[0]
	metrics.MemoryUsage = ramPercentage.UsedPercent
	metrics.DiskUsage = diskUsage.UsedPercent

	return &metrics, nil
}
