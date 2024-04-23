package utils

import (
	"os"
	"log"
	"time"
)

// Logger represents a configurable logger
type Logger struct {
	*log.Logger
	MaxFileSize int64
	MaxBackups  int
}

// NewLogger creates a new logger with specified filename, max file size, and max backups
func NewLogger(filename string, maxFileSize int64, maxBackups int) *Logger {
	// Create log file
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Failed to open log file: %s", err)
	}

	// Initialize logger with custom settings
	logger := log.New(file, "", log.Ldate|log.Ltime)
	l := &Logger{
		Logger:      logger,
		MaxFileSize: maxFileSize,
		MaxBackups:  maxBackups,
	}

	// Start log rotation routine
	go l.rotateLogs(filename)

	return l
}

// rotateLogs rotates log files based on size and number of backups
func (l *Logger) rotateLogs(filename string) {
	for {
		time.Sleep(24 * time.Hour) // Rotate logs daily
		info, err := os.Stat(filename)
		if err != nil {
			log.Printf("Failed to rotate logs: %s", err)
			continue
		}

		fileSize := info.Size()
		if fileSize >= l.MaxFileSize {
			err = os.Rename(filename, filename+".bak")
			if err != nil {
				log.Printf("Failed to rotate logs: %s", err)
				continue
			}
			// Delete old backups if exceeds max number of backups
			backups, err := filepath.Glob(filename + ".*.bak")
			if err != nil {
				log.Printf("Failed to rotate logs: %s", err)
				continue
			}
			if len(backups) > l.MaxBackups {
				sort.Strings(backups)
				for i := 0; i < len(backups)-l.MaxBackups; i++ {
					err := os.Remove(backups[i])
					if err != nil {
						log.Printf("Failed to rotate logs: %s", err)
						continue
					}
				}
			}
		}
	}
}
