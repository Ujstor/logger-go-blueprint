package logger

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
	"time"
)

var (
	fileLogger *log.Logger
	stdLogger  *log.Logger
)

func init() {
	logFilePath := filepath.Join(".", "log.json")
	file, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening log file: %v", err)
	}

	fileLogger = log.New(file, "", 0)

	stdLogger = log.New(os.Stdout, "", 0)
}

type logEntry struct {
	Time       string      `json:"time"`
	Level      string      `json:"level"`
	Message    string      `json:"message"`
	Attributes interface{} `json:"attributes,omitempty"`
}

func logJSON(logger *log.Logger, level string, msg string, attrs ...interface{}) {
	entry := logEntry{
		Time:    time.Now().Format("2006/01/02 15:04:05"),
		Level:   level,
		Message: msg,
	}
	if len(attrs) > 0 {
		entry.Attributes = attrs
	}
	data, err := json.Marshal(entry)
	if err != nil {
		log.Printf("error marshalling log entry: %v", err)
		return
	}
	data = append(data, '\n')
	logger.Println(string(data))
}

func Info(msg string, attrs ...interface{}) {

}

func Middleware(msg string, attrs ...interface{}) {
	logJSON(fileLogger, "MIDDLEWARE", msg, attrs...)
	logJSON(stdLogger, "MIDDLEWARE", msg, attrs...)
}

func Warn(msg string, attrs ...interface{}) {
	logJSON(fileLogger, "WARN", msg, attrs...)
	logJSON(stdLogger, "WARN", msg, attrs...)
}

func Error(msg string, attrs ...interface{}) {
	logJSON(fileLogger, "ERROR", msg, attrs...)
	logJSON(stdLogger, "ERROR", msg, attrs...)
}
