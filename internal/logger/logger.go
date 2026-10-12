package logger

import "github.com/sirupsen/logrus"

var log = logrus.New()

func Debug(args ...any)                 { log.Debug(args...) }
func Debugf(format string, args ...any) { log.Debugf(format, args...) }
func Info(args ...any)                  { log.Info(args...) }
func Infof(format string, args ...any)  { log.Infof(format, args...) }
func Warn(args ...any)                  { log.Warn(args...) }
func Warnf(format string, args ...any)  { log.Warnf(format, args...) }
func Error(args ...any)                 { log.Error(args...) }
func Errorf(format string, args ...any) { log.Errorf(format, args...) }
func Fatal(args ...any)                 { log.Fatal(args...) }
func Fatalf(format string, args ...any) { log.Fatalf(format, args...) }

func Init(isDebug bool) {
	log.Formatter = &logrus.TextFormatter{
		FullTimestamp: true,
	}
	log.Level = logrus.InfoLevel

	if isDebug {
		log.Level = logrus.DebugLevel
	}
}
