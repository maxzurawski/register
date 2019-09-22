package config

import (
	"fmt"
	"os"

	"github.com/xdevices/utilities/config"
)

type registerConfig struct {
	config.Manager
	dbPath string
}

var instance *registerConfig

func Config() *registerConfig {
	if instance == nil {
		instance = new(registerConfig)
		instance.Init()
		instance.registerConfigInit()
	}
	return instance
}

func (c *registerConfig) registerConfigInit() {
	if dbPath, err := os.LookupEnv("DB_PATH"); !err {
		panic(fmt.Sprintf("set DB_PATH and try again"))
	} else {
		c.dbPath = dbPath
	}
}

func (c *registerConfig) DBPath() string {
	return c.dbPath
}
