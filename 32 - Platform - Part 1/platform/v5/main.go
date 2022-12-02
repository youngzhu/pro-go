package main

import (
	"platform/v5/config"
	"platform/v5/logging"
	"platform/v5/services"
)

func writeMessage(logger logging.Logger, cfg config.Configuration) {
	section, ok := cfg.GetSection("main")
	if ok {
		msg, ok := section.GetString("message")
		if ok {
			logger.Info(msg)
		} else {
			logger.Panic("Cannot find configuration setting")
		}
	} else {
		logger.Panic("Config section not found")
	}
}

func main() {
	services.RegisterDefaultServices()

	services.Call(writeMessage)

	val := struct {
		message string
		logging.Logger
	}{
		message: "Hello from the struct",
	}
	services.Populate(&val)
	val.Logger.Debug(val.message)
}
