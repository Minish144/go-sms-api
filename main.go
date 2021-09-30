package main

import (
	"github.com/minish144/go-sms-api/modules/server"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	// viper init
	viper.AddConfigPath("./")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Fatalln()
	}
	logrus.Info("viper config initialized successfully")

	// server init
	server.Run()
}
