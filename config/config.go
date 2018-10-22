package config

import (
	"sync"
	"os"
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
			ENV: os.Getenv("ENV"),
			DB: &DBConfig{
				Dialect:  "mysql",
				Username: os.Getenv("DBUSERNAME"),
				Password: os.Getenv("DBPASSWORD"),
				Name:     os.Getenv("DBNAME"),
				Charset:  "utf8",
			},
		}
	})

	return config;
}