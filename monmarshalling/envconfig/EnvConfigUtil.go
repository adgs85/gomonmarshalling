package envconfig

import (
	"log"

	"github.com/spf13/viper"
)

var config *viper.Viper = initViperConfig()

func initViperConfig() *viper.Viper {
	viperInst := viper.New()
	viperInst.AddConfigPath(".")
	viperInst.AutomaticEnv()

	viperError := viperInst.ReadInConfig()

	if viperError != nil {
		log.Fatalf("%v", viperError)
	}

	return viperInst
}

func GetViperConfig() *viper.Viper {
	return config
}
