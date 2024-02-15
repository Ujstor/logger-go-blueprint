package main

import (
	"logger-test/logger" 
	"logger-test/internal/server"
	"os"
)

func main() {
	server := server.NewServer()

	err := server.ListenAndServe()
	if err != nil {
		logger.Warn("Cannot start server", "error", err)
		os.Exit(1)
	}
}
