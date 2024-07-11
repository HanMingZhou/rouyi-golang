package ruoyilog

import "github.com/sirupsen/logrus"

var logRu *logrus.Logger

func GetLog() *logrus.Logger {
	if logRu == nil {
		logRu = logrus.New()
	}
	return logRu
}
func Warn(args ...interface{}) {
	GetLog().Warn(args)
}
