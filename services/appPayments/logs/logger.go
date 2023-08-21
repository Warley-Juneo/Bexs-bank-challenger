package logs

import (
	"fmt"
	"log"
	"os"
)

type JsonResponseClient struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func LogToFile(filePath, message string) error {
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		return fmt.Errorf("failed to open log file: %w", err)
	}
	defer file.Close()

	logger := log.New(file, "", log.LstdFlags)
	logger.Println(message)

	return nil
}
