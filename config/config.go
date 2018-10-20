package config

import (
	"sync"
)

type DBConfig struct {
	Dialect  string
	Username string
	Password string
	Name     string
	Charset  string
}

type Config struct {
	ENV string
	DB *DBConfig
}

var config *Config
var once sync.Once

func GetConfig() *Config {
	once.Do(func() {
		config = &Config{
			ENV: "",
			DB: &DBConfig{
				Dialect:  "mysql",
				Username: "root",
				Password: "root",
				Name:     "todoapp",
				Charset:  "utf8",
			},
		}
	})

	return config;
}

func (config *Config) SetEnv(env string) {
	config.ENV = env;
}