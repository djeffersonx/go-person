package config

import (
	"errors"
	"fmt"
	"log"
	"os"
)

var (
	ApplicationName           string
	ApplicationConfigs        []string
	applicationConfigProvider string
	applicationProfile        string
	providers                 = map[string]ApplicationConfigurationProvider{
		SSMConfigProviderName:    ssmConfigProvider{},
		StaticConfigProviderName: staticConfigProvider{},
	}
)

const (
	ParamApplicationProfile        = "application.profile"
	ParamApplicationConfigProvider = "application.configProvider"
	SSMConfigProviderName          = "ssm_config_provider"
	StaticConfigProviderName       = "static_config_provider"
)

func Init() {

	applicationConfigProvider = os.Getenv(ParamApplicationConfigProvider)
	applicationProfile = os.Getenv(ParamApplicationProfile)

	log.Println("Starting configuration ...")

	validateRequiredConfig(ParamApplicationProfile, applicationProfile)
	validateRequiredConfig(ParamApplicationConfigProvider, applicationProfile)

	providers[applicationConfigProvider].LoadConfiguration()

}

func validateRequiredConfig(configName string, configValue string) {
	if configValue == "" {
		panic(errors.New(fmt.Sprintf("Configuration: %s is required", configName)))
	}
}

type ApplicationConfigurationProvider interface {
	LoadConfiguration()
}
