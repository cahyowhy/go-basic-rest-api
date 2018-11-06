package cloudinary

import (
	"go-basic-rest-api/config"
	"log"
	"sync"
)

type singleton struct {
}

var instance *Service
var once sync.Once

func GetService() *Service {
	config := config.GetConfig()
	var accountKey = config.CloudinaryConfig.Acount
	var secretKey = config.CloudinaryConfig.Secret
	var cloudName = config.CloudinaryConfig.CloudName

	endpoint := "cloudinary://" + accountKey + ":" + secretKey + "@" + cloudName
	once.Do(func() {
		var err error
		instance, err = Dial(endpoint)
		if err != nil {
			log.Fatal(err)
		}
	})
	return instance
}
