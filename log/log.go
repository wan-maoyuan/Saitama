package log

import (
	"github.com/sirupsen/logrus"
)

type Logger interface {
	Log(level Level, keyvals ...any) error
}

type logger struct {
	logger *logrus.Logger
}

func Default() *logger {
	l := logrus.New()

	return &logger{
		logger: l,
	}
}

func (l *logger) Log(level Level, keyvals ...any) (err error) {
	var logrusLevel logrus.Level

	switch level {
	case LevelDebug:
		logrusLevel = logrus.DebugLevel
	case LevelInfo:
		logrusLevel = logrus.InfoLevel
	case LevelWarn:
		logrusLevel = logrus.WarnLevel
	case LevelError:
		logrusLevel = logrus.ErrorLevel
	case LevelFatal:
		logrusLevel = logrus.FatalLevel
	default:
		logrusLevel = logrus.DebugLevel
	}

	l.logger.Log(logrusLevel, keyvals)
	return
}
