package main

import (
	"platform/v2/config"
	"platform/v2/logging"
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
	cfg, err := config.Load("config.json")
	if err != nil {
		panic(err)
	}
	logger := logging.NewDefaultLogger(cfg)
	writeMessage(logger, cfg)
}
