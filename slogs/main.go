package main

import (
	"fmt"
	"log/slog"
	"os"
	"time"
)

const logDir = "logs"
const logFileFormat = "2006-01-02"

func InitLogger() (*slog.Logger, error) {

	if _, err := os.Stat(logDir); os.IsNotExist(err) {
		err := os.Mkdir(logDir, 0755)
		if err != nil {
			return nil, fmt.Errorf("failed to create log directory: %v", err)
		}
	}

	logFileName := fmt.Sprintf("%s/%s.log", logDir, time.Now().Format(logFileFormat))
	logFile, err := os.OpenFile(logFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, fmt.Errorf("failed to open log file: %v", err)
	}

	logger := slog.New(slog.NewJSONHandler(logFile, nil))

	return logger, nil
}

func main() {
	logger, err := InitLogger()
	if err != nil {
		fmt.Println("Error initializing logger:", err)
		return
	}
	
	for i := 0; i < 999999; i++ {
		logger.Info("Logger initialized successfully")
		i++;
	}

}