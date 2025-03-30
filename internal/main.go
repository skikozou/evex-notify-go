package main

import (
	"evex-notify/config"

	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetLevel(logrus.ErrorLevel)
	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors:            true,
		DisableLevelTruncation: true,
		PadLevelText:           true,
	})

	_, err := config.LoadConfig()
	if err != nil {
		logrus.Error(err)
	}
}
