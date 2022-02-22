package config

import (
	"go-person/pkg/config"
	"os"
)

const databaseUrl = "database.url"
const databaseUsername = "database.username"
const databasePassword = "database.password"

func init() {
	setApplicationName()
	setApplicationConfigKeys()

	// <temporaryBlock> this is for initial tests only all will be served by env vars...
	setApplicationProfile()
	setConfigProvider()
	// </temporaryBlock>

	config.EnvironmentConfigs = config.EnvironmentConfig{"dev": devEnvironment}

	config.Init()

}

func setConfigProvider() error {
	return os.Setenv(config.ParamApplicationConfigProvider, config.StaticConfigProviderName)
}

func setApplicationProfile() error {
	return os.Setenv(config.ParamApplicationProfile, "dev")
}

func setApplicationName() {
	config.ApplicationName = "go-person"
}

func setApplicationConfigKeys() {
	config.ApplicationConfigs = []string{
		databaseUrl,
		databaseUsername,
		databasePassword,
	}
}

var (
	devEnvironment = map[string]string{
		databaseUrl:      "",
		databaseUsername: "",
		databasePassword: "",
	}
)
