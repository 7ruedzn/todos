package config

import (
	"fmt"
	"log/slog"
	"os"
)

var (
	LogFile *os.File
)

func SetupLogs() error {
	if AppInstance.Config.LogsPath == "" {
		return fmt.Errorf("path to logs file cannot be empty")
	}

	LogFile, err := os.OpenFile(AppInstance.Config.LogsPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)

	if err != nil {
		return err
	}

	handler := slog.NewJSONHandler(LogFile, &slog.HandlerOptions{
		AddSource: true,
	})
	slogger := slog.New(handler)
	slog.SetDefault(slogger)

	return nil
}
