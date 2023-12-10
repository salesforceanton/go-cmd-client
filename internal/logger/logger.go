package logger

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
)

func SetConfiguration() {
	logrus.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "02-01-2006 15:04:05",
		DisableQuote:    true,
	})
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.InfoLevel)
}

func LogInfo(point, info string) {
	msg := fmt.Sprintf("\n[POINT]: %s [INFO]: %s", point, info)
	logrus.Info(msg)
}

func LogError(point string, err error) {
	msg := fmt.Sprintf("\n[POINT]: %s [ERROR]: %s", point, err)
	logrus.Error(msg)
}
