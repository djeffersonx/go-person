package config

import "go-person/pkg/config"

func init() {
	config.ApplicationName = "go-person"
	config.ApplicationConfigs = append(config.ApplicationConfigs, "database.url")
	config.ApplicationConfigs = append(config.ApplicationConfigs, "database.username")
	config.ApplicationConfigs = append(config.ApplicationConfigs, "database.password")

	config.ApplicationConfigProvider = config.StaticConfigProvider
}
