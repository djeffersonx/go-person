package config

import (
	"errors"
	"fmt"
	"os"
)

type ApplicationConfigKeys []string

var (
	ApplicationConfigProvider = os.Getenv("application.configProvider")
	ApplicationName           = ""
	ApplicationConfigs        = ApplicationConfigKeys{}
	ApplicationProfile        = os.Getenv("application.profile")
	Providers                 = map[string]ApplicationConfigurationProvider{
		SSMConfigProviderName:    SSMConfigProvider{},
		StaticConfigProviderName: StaticConfigProvider{},
	}
)

const (
	SSMConfigProviderName    = "ssmConfigProvider"
	StaticConfigProviderName = "staticConfigProvider"
)

func init() {

	validateRequiredConfig("applicationProfile", ApplicationProfile)
	setDefault(ApplicationConfigProvider, StaticConfigProviderName)

	Providers[ApplicationConfigProvider].LoadConfiguration()

}

func setDefault(_var string, defaultValue string) {
	if _var == "" {
		_var = defaultValue
	}
}

func validateRequiredConfig(configName string, configValue string) {
	if configValue == "" {
		panic(errors.New(fmt.Sprintf("Configuraiton: %s is required", configName)))
	}
}

type ApplicationConfigurationProvider interface {
	LoadConfiguration()
}
