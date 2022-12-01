package main

import (
	"platform/v1/logging"
)

func writeMessage(logger logging.Logger) {
	logger.Info("Hello Platform")
}

func main() {
	logger := logging.NewDefaultLogger(logging.Information)
	writeMessage(logger)
}
