package helpers

import (
	"log"
	"os"
	"time"
)

const maxLogFileSize = 1 * 1024 * 1024 // 1 MB

func WriteLog(filePath string, logID, logMessage string) {
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal("Failed to open log file: ", err)
	}
	defer file.Close()

	// Create a new logger that writes to the log file
	logger := log.New(file, logID, log.Ldate|log.Ltime|log.Lshortfile)

	// Write some log messages
	logger.Println(logMessage)
}

func RotateTruncateLog(path string, maxFileSize int64) {
	logFile, err := os.OpenFile("app.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Error opening log file: %v", err)
	}
	defer logFile.Close()

	log.SetOutput(logFile)

	for {
		info, err := logFile.Stat()
		if err != nil {
			log.Printf("Error getting log file stats: %v", err)
		}

		if info.Size() >= maxLogFileSize {
			err = truncateLogFile(logFile)
			if err != nil {
				log.Printf("Error truncating log file: %v", err)
			}
		}

		time.Sleep(time.Minute)
	}
}

func truncateLogFile(file *os.File) error {
	err := file.Truncate(0)
	if err != nil {
		return err
	}

	_, err = file.Seek(0, 0)
	if err != nil {
		return err
	}

	return nil
}
