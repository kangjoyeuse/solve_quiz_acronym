package internal

import (
	"fmt"
	"log/slog"
)

type Monitor struct{}

func NewMonitor() *Monitor {
	return &Monitor{}
}

func (m *Monitor) RefreshDisplay() error {
	fmt.Print("\033[H\033[2J")

	cpuUsage, err := GetCPUUsage()
	if err != nil {
		slog.Error("failed to get CPU usage", "error", err)
		cpuUsage = 0.0
	}
	fmt.Printf("CPU Usage: %.2f%%\n", cpuUsage)

	memTotal, memFree, err := GetMemoryUsage()
	if err != nil {
		slog.Error("failed to get memory usage", "error", err)
		memTotal, memFree = 0, 0
	}
	fmt.Printf("Memory: %d MB total, %d MB free\n", memTotal, memFree)

	diskTotal, diskFree, err := GetDiskUsage("/")
	if err != nil {
		slog.Error("failed to get disk usage", "error", err)
		diskTotal, diskFree = 0, 0
	}
	fmt.Printf("Disk: %d GB total, %d GB free\n", diskTotal, diskFree)

	procCount, err := CountProcesses()
	if err != nil {
		slog.Error("failed to count processes", "error", err)
		procCount = 0
	}
	fmt.Printf("Running Processes: %d\n", procCount)

	fmt.Println("\nPress Ctrl+C to exit...")
	return nil
}
