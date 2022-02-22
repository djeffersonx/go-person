package config

import (
	"os"
)

type staticConfigProvider struct{}

type EnvironmentConfig map[string]map[string]string

var (
	EnvironmentConfigs EnvironmentConfig
)

func (configProvide staticConfigProvider) LoadConfiguration() {
	for _, configKey := range ApplicationConfigs {
		os.Setenv(configKey, EnvironmentConfigs[applicationProfile][configKey])
	}
}
