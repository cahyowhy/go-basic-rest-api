package config

import (
	"os"
	"sync"
)

type DBConfig struct {
	Dialect  string
	Username string
	Password string
	Name     string
	Charset  string
	Host     string
	Port     string
}

type Config struct {
	ENV string
	DB  *DBConfig
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
				Host:     os.Getenv("DBHOST"),
				Port:     os.Getenv("DBPORT"),
				Charset:  "utf8",
			},
		}
	})

	return config
}
