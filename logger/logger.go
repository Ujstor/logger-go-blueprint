package logger

import (
	"encoding/json"
	"log"
	"os"
	"time"
)

var (
	fileLogger *log.Logger
	stdLogger  *log.Logger
)

func init() {
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
	logJSON(stdLogger, "INFO", msg, attrs...)

}

func Middleware(msg string, attrs ...interface{}) {
	logJSON(stdLogger, "MIDDLEWARE", msg, attrs...)
}

func Warn(msg string, attrs ...interface{}) {
	logJSON(stdLogger, "WARN", msg, attrs...)
}

func Error(msg string, attrs ...interface{}) {
	logJSON(stdLogger, "ERROR", msg, attrs...)
}
