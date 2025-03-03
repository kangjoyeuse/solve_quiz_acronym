package internal

import (
	"fmt"
	"log/slog"
	"os"
	"strconv"
	"strings"
)

func GetMemoryUsage() (uint64, uint64, error) {
	data, err := os.ReadFile("/proc/meminfo")
	if err != nil {
		slog.Error("error reading /proc/meminfo", "error", err)
		return 0, 0, fmt.Errorf("read /proc/meminfo: %w", err)
	}
	
	memInfo := make(map[string]uint64)
	lines := strings.Split(string(data), "\n")
	
	for _, line := range lines {
		fields := strings.Fields(line)
		if len(fields) < 2 {
			continue
		}
		
		key := strings.TrimSuffix(fields[0], ":")
		val, err := strconv.ParseUint(fields[1], 10, 64)
		if err != nil {
			slog.Debug("failed to parse memory value", "key", key, "value", fields[1], "error", err)
			continue
		}
		
		memInfo[key] = val
	}
	
	total, ok1 := memInfo["MemTotal"]
	free, ok2 := memInfo["MemFree"]
	
	if !ok1 || !ok2 {
		return 0, 0, fmt.Errorf("required memory info fields not found")
	}
	
	return total / 1024, free / 1024, nil
}