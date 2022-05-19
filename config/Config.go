package config

import (
	"fmt"
	"os"
	"sync"
)

type AppConfig struct {
	Port     int
	Driver   string
	Name     string
	Address  string
	DB_Port  int
	Username string
	Password string
}

var lock = &sync.Mutex{}
var appConfig *AppConfig

func GetConfig() *AppConfig {
	lock.Lock()
	defer lock.Unlock()

	if appConfig == nil {
		appConfig = InitConfig()
	}

	return appConfig
}

func InitConfig() *AppConfig {
	var defaultConfig AppConfig
	defaultConfig.Port = 8000
	defaultConfig.Driver = GetEnv("DRIVER", "mysql")
	defaultConfig.Name = GetEnv("NAME", "event")
	defaultConfig.Address = GetEnv("ADDRESS", "localhost")
	defaultConfig.DB_Port = 3306
	defaultConfig.Username = GetEnv("DB_USERNAME", "root")
	defaultConfig.Password = GetEnv("PASSWORD", "")

	fmt.Println(defaultConfig)

	return &defaultConfig
}

func GetEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		fmt.Println(value)
		return value
	}

	return fallback
}
