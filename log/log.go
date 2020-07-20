package log

import "github.com/sirupsen/logrus"

var logger *logrus.Logger

func Init(level logrus.Level) (err error) {
	logger = logrus.New()
	logger.SetLevel(level)

	return
}

func Trace(fields logrus.Fields, args ...interface{}) {
	logger.WithFields(fields).Trace(args)
}

func Debug(fields logrus.Fields, args ...interface{}) {
	logger.WithFields(fields).Debug(args)
}

func Info(fields logrus.Fields, args ...interface{}) {
	logger.WithFields(fields).Info(args)
}

func Warn(fields logrus.Fields, args ...interface{}) {
	logger.WithFields(fields).Warn(args)
}

func Error(fields logrus.Fields, args ...interface{}) {
	logger.WithFields(fields).Error(args)
}
