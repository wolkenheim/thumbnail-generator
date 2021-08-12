package app

import (
"fmt"
"github.com/spf13/viper"
)

const EnvLocal string = "local"
const EnvTesting string = "testing"

func InitConfig() {
	viper.BindEnv("APP_ENV", "APP_ENV")
	viper.SetDefault("APP_ENV", EnvLocal)
	readConfig(viper.GetString( "APP_ENV"))
}

func readConfig(envName string) {
	viper.SetConfigName(envName)
	viper.SetConfigType("yaml")

	if envName == EnvLocal || envName == EnvTesting {
		viper.AddConfigPath("./config")
	} else {
		viper.BindEnv("minio.secretAccessKey", "MINIO_SECRET")
		viper.BindEnv("dam.password", "DAM_ELVIS_PASSWORD")
		viper.AddConfigPath("/app/config")
	}

	viper.SetDefault("server.port", ":3001")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}

