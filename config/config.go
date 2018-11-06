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

type CloudinaryConfig struct {
	Acount    string
	Secret    string
	CloudName string
}

type Config struct {
	ENV              string
	DB               *DBConfig
	CloudinaryConfig *CloudinaryConfig
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
			CloudinaryConfig: &CloudinaryConfig{
				Acount:    os.Getenv("Cloudinary_Account_Key"),
				Secret:    os.Getenv("Cloudinary_Secret_Key"),
				CloudName: os.Getenv("Cloudinary_Cloud_Name"),
			},
		}
	})

	return config
}
