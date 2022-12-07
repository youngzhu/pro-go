package main

import (
	"platform/logging"
	"platform/services"
)

func writeMessage(logger logging.Logger) {
	logger.Info("Sports Store")
}

func main() {
	services.RegisterDefaultServices()
	services.Call(writeMessage)
}
