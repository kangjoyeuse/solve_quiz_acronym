package internal

import (
	"fmt"
	"log/slog"
	"os"
	"strconv"
	"strings"
)

type CPUStats struct {
	User   float64
	Nice   float64
	System float64
	Idle   float64
	Total  float64
}

func GetCPUUsage() (float64, error) {
	data, err := os.ReadFile("/proc/stat")
	if err != nil {
		slog.Error("error reading /proc/stat", "error", err)
		return 0, fmt.Errorf("read /proc/stat: %w", err)
	}
	
	lines := strings.Split(string(data), "\n")
	if len(lines) == 0 {
		return 0, fmt.Errorf("empty /proc/stat")
	}
	
	fields := strings.Fields(lines[0])
	if len(fields) < 5 || fields[0] != "cpu" {
		return 0, fmt.Errorf("invalid /proc/stat format")
	}
	
	stats, err := parseCPUStats(fields)
	if err != nil {
		return 0, err
	}
	
	usage := 100.0 * (stats.User + stats.Nice + stats.System) / stats.Total
	return usage, nil
}

func parseCPUStats(fields []string) (CPUStats, error) {
	var stats CPUStats
	var err error
	
	if stats.User, err = strconv.ParseFloat(fields[1], 64); err != nil {
		return stats, fmt.Errorf("parse user cpu: %w", err)
	}
	
	if stats.Nice, err = strconv.ParseFloat(fields[2], 64); err != nil {
		return stats, fmt.Errorf("parse nice cpu: %w", err)
	}
	
	if stats.System, err = strconv.ParseFloat(fields[3], 64); err != nil {
		return stats, fmt.Errorf("parse system cpu: %w", err)
	}
	
	if stats.Idle, err = strconv.ParseFloat(fields[4], 64); err != nil {
		return stats, fmt.Errorf("parse idle cpu: %w", err)
	}
	
	stats.Total = stats.User + stats.Nice + stats.System + stats.Idle
	return stats, nil
}