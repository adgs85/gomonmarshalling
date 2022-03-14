package envconfig

import (
	"log"

	"github.com/spf13/viper"
)

var config *viper.Viper

func initViperConfig() {
	if config == nil {
		viperInst := viper.New()
		viperInst.AddConfigPath(".")
		viperInst.AutomaticEnv()

		viperError := viperInst.ReadInConfig()

		if viperError != nil {
			log.Fatalf("%v", viperError)
		}
		config = viperInst
	}
}

func GetViperConfig() *viper.Viper {
	initViperConfig()
	return config
}
