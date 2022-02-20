package config

import "go-person/pkg/config"

func init() {
	config.ApplicationName = "go-person"
	config.Configuration["database.url"] = ""
	config.Configuration["database.username"] = ""
	config.Configuration["database.password"] = ""
}
