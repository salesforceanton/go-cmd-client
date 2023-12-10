package logger

import (
	"github.com/sirupsen/logrus"
)

func LogInfo(point, info string) {
	logrus.WithFields(logrus.Fields{
		"POINT": point,
		"INFO:": info,
	}).Info()
}

func LogError(point string, err error) {
	logrus.WithFields(logrus.Fields{
		"POINT": point,
		"ERROR": err.Error(),
	}).Error()
}
