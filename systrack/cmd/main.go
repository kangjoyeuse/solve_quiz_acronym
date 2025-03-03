package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"systrack/internal"
	"time"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}))
	slog.SetDefault(logger)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	go func() {
		s := <-sigChan
		slog.Info("shutdown signal received", "signal", s)
		fmt.Println("\nExiting SysTrack. Goodbye!")
		cancel()
		os.Exit(0)
	}()

	fmt.Println("SysTrack - Minimal System Monitor (Press Ctrl+C to exit)")

	mn := internal.NewMonitor()
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			if err := mn.RefreshDisplay(); err != nil {
				slog.Error("error refreshing display", "error", err)
			}
		}
	}
}
