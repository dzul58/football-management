package logger

import (
	"log"
	"os"
)

var (
	InfoLogger  *log.Logger
	ErrorLogger *log.Logger
	WarnLogger  *log.Logger
)

// InitLogger initializes loggers
func InitLogger(logFile string) error {
	// Create logs directory if not exists
	if err := os.MkdirAll("logs", 0755); err != nil {
		return err
	}

	// Open log file
	file, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}

	InfoLogger = log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	WarnLogger = log.New(file, "WARN: ", log.Ldate|log.Ltime|log.Lshortfile)

	return nil
}

// Info logs info messages
func Info(message string) {
	if InfoLogger != nil {
		InfoLogger.Println(message)
	}
	log.Println("INFO:", message)
}

// Error logs error messages
func Error(message string) {
	if ErrorLogger != nil {
		ErrorLogger.Println(message)
	}
	log.Println("ERROR:", message)
}

// Warn logs warning messages
func Warn(message string) {
	if WarnLogger != nil {
		WarnLogger.Println(message)
	}
	log.Println("WARN:", message)
}
