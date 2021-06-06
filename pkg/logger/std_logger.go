package logger

import (
	"fmt"
	"os"
	"sync"
)

type stdLogger struct {
	logLevel       LogLevel
	customPrefixFn func() string
	mux            sync.Mutex // NOTE: all logging to fmt is actually NOT thread safe.
}

func emptyPrefixFn() string { return "" }

func NewStdLogger(logLevel LogLevel) Logger {
	l := &stdLogger{logLevel: logLevel, customPrefixFn: emptyPrefixFn}
	return l
}

func (l *stdLogger) SetPrefix(prefixFn func() string) {
	l.mux.Lock()
	defer l.mux.Unlock()
	l.customPrefixFn = prefixFn
}

func (l *stdLogger) Level() LogLevel { return l.logLevel }
func (l *stdLogger) SetLevel(lvl LogLevel) {
	l.mux.Lock()
	defer l.mux.Unlock()
	l.logLevel = lvl
}

func (l *stdLogger) Flush() error {
	l.Infoln("Flushing logger")
	return nil
}

func (l *stdLogger) Debug(v ...interface{}) {
	l.mux.Lock()
	defer l.mux.Unlock()
	if l.Level() <= DebugLevel {
		prefix := l.customPrefixFn()
		fmt.Printf("[DEBUG] %s", prefix)
		fmt.Print(v...)
	}
}
func (l *stdLogger) Debugf(format string, v ...interface{}) {
	l.mux.Lock()
	defer l.mux.Unlock()
	if l.Level() <= DebugLevel {
		prefix := l.customPrefixFn()
		fmt.Printf("[DEBUG] %s", prefix)
		fmt.Printf(format, v...)
	}
}
func (l *stdLogger) Debugln(v ...interface{}) {
	l.mux.Lock()
	defer l.mux.Unlock()
	if l.Level() <= DebugLevel {
		prefix := l.customPrefixFn()
		fmt.Printf("[DEBUG] %s", prefix)
		fmt.Println(v...)
	}
}

func (l *stdLogger) Info(v ...interface{}) {
	l.mux.Lock()
	defer l.mux.Unlock()
	if l.Level() <= InfoLevel {
		prefix := l.customPrefixFn()
		fmt.Printf("[INFO] %s", prefix)
		fmt.Print(v...)
	}
}
func (l *stdLogger) Infof(format string, v ...interface{}) {
	l.mux.Lock()
	defer l.mux.Unlock()
	if l.Level() <= InfoLevel {
		prefix := l.customPrefixFn()
		fmt.Printf("[INFO] %s", prefix)
		fmt.Printf(format, v...)
	}
}
func (l *stdLogger) Infoln(v ...interface{}) {
	l.mux.Lock()
	defer l.mux.Unlock()
	if l.Level() <= InfoLevel {
		prefix := l.customPrefixFn()
		fmt.Printf("[INFO] %s", prefix)
		fmt.Println(v...)
	}
}

func (l *stdLogger) Warn(v ...interface{}) {
	l.mux.Lock()
	defer l.mux.Unlock()
	if l.Level() <= WarnLevel {
		prefix := l.customPrefixFn()
		fmt.Printf("[WARN] %s", prefix)
		fmt.Print(v...)
	}
}
func (l *stdLogger) Warnf(format string, v ...interface{}) {
	l.mux.Lock()
	defer l.mux.Unlock()
	if l.Level() <= WarnLevel {
		prefix := l.customPrefixFn()
		fmt.Printf("[WARN] %s", prefix)
		fmt.Printf(format, v...)
	}
}
func (l *stdLogger) Warnln(v ...interface{}) {
	l.mux.Lock()
	defer l.mux.Unlock()
	if l.Level() <= WarnLevel {
		prefix := l.customPrefixFn()
		fmt.Printf("[WARN] %s", prefix)
		fmt.Println(v...)
	}
}

func (l *stdLogger) Error(v ...interface{}) {
	l.mux.Lock()
	defer l.mux.Unlock()
	if l.Level() <= ErrorLevel {
		prefix := l.customPrefixFn()
		fmt.Printf("[ERROR] %s", prefix)
		fmt.Fprint(os.Stderr, v...)
	}
}
func (l *stdLogger) Errorf(format string, v ...interface{}) {
	l.mux.Lock()
	defer l.mux.Unlock()
	if l.Level() <= ErrorLevel {
		prefix := l.customPrefixFn()
		fmt.Printf("[ERROR] %s", prefix)
		fmt.Fprintf(os.Stderr, format, v...)
	}
}
func (l *stdLogger) Errorln(v ...interface{}) {
	l.mux.Lock()
	defer l.mux.Unlock()
	if l.Level() <= ErrorLevel {
		prefix := l.customPrefixFn()
		fmt.Printf("[ERROR] %s", prefix)
		fmt.Fprintln(os.Stderr, v...)
	}
}
