package services

import (
	"platform/config"
	"platform/logging"
	"platform/services"
	"platform/templates"
	"platform/validation"
)

func RegisterServices() {

	err := services.AddSingleton(func() (c config.Configuration) {
		c, loadErr := config.Load("v2/config.json")
		if loadErr != nil {
			panic(loadErr)
		}
		return
	})

	err = services.AddSingleton(func(appconfig config.Configuration) logging.Logger {
		return logging.NewDefaultLogger(appconfig)
	})
	if err != nil {
		panic(err)
	}

	err = services.AddSingleton(
		func(c config.Configuration) templates.TemplateExecutor {
			templates.LoadTemplates(c)
			return &templates.LayoutTemplateProcessor{}
		})
	if err != nil {
		panic(err)
	}

	err = services.AddSingleton(
		func() validation.Validator {
			return validation.NewDefaultValidator(validation.DefaultValidators())
		})
	if err != nil {
		panic(err)
	}

}
