package main

import (
	"platform/v4/config"
	"platform/v4/logging"
	"platform/v4/services"
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
}
