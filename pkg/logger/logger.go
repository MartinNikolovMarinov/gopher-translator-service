package logger

type LogLevel uint8

const (
	DebugLevel LogLevel = 0
	InfoLevel  LogLevel = 1
	WarnLevel  LogLevel = 2
	ErrorLevel LogLevel = 3
)

type Logger interface {
	SetPrefix(func() string)

	Level() LogLevel
	SetLevel(lvl LogLevel)

	Debug(v ...interface{})
	Debugf(format string, v ...interface{})
	Debugln(v ...interface{})

	Info(v ...interface{})
	Infof(format string, v ...interface{})
	Infoln(v ...interface{})

	Warn(v ...interface{})
	Warnf(format string, v ...interface{})
	Warnln(v ...interface{})

	Error(v ...interface{})
	Errorf(format string, v ...interface{})
	Errorln(v ...interface{})

	Flush() error
}
