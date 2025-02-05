package config

import (
	"fmt"
	"log"
	"os"
)

var (
	InfoLog    *log.Logger
	WarningLog *log.Logger
	ErrorLog   *log.Logger
	LogFile    *os.File
)

func SetupLogs() error {
	if AppInstance.Config.LogsPath == "" {
		return fmt.Errorf("path to logs file cannot be empty")
	}

	LogFile, err := os.OpenFile(AppInstance.Config.LogsPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)

	if err != nil {
		return err
	}

	InfoLog = log.New(LogFile, "INFO: ", log.LstdFlags|log.Lshortfile)
	WarningLog = log.New(LogFile, "WARNING: ", log.LstdFlags|log.Lshortfile)
	ErrorLog = log.New(LogFile, "ERROR: ", log.LstdFlags|log.Lshortfile)

	return nil
}
