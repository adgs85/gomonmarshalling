package envconfig

import (
	"log"

	"github.com/spf13/viper"
)

var config *viper.Viper

func initViperConfig() *viper.Viper {
	viperInst := viper.New()
	viperInst.AddConfigPath(".")
	viperInst.AutomaticEnv()

	viperError := viperInst.ReadInConfig()

	if viperError != nil {
		log.Panicf("%v", viperError)
	}

	return viperInst
}

func GetViperConfig() *viper.Viper {
	if config == nil {
		config = initViperConfig()
	}
	return config
}
