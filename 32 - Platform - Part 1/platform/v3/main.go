package main

import (
	"platform/v3/config"
	"platform/v3/logging"
	"platform/v3/services"
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

	var cfg config.Configuration
	services.GetService(&cfg)

	var logger logging.Logger
	services.GetService(&logger)

	writeMessage(logger, cfg)
}
