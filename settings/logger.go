package settings

import (
	"io"
	"log"
	"os"

	"github.com/fatih/color"
)

type Logger struct {
	debug   *log.Logger
	info    *log.Logger
	warning *log.Logger
	err     *log.Logger
	writer  io.Writer
}

// Logger
func NewLogger(p string) *Logger {
	writer := io.Writer(os.Stdout)
	logger := log.New(writer, p, log.Ldate|log.Ltime)

	return &Logger{
		debug:   log.New(writer, color.New(color.FgBlue).SprintFunc()("DEBUG "), logger.Flags()),
		info:    log.New(writer, color.New(color.FgGreen).SprintFunc()("INFO "), logger.Flags()),
		warning: log.New(writer, color.New(color.FgYellow).SprintFunc()("WARNING "), logger.Flags()),
		err:     log.New(writer, color.New(color.FgRed).SprintFunc()("ERROR "), logger.Flags()),
		writer:  writer,
	}
}

// Create Non-Formatted Logs
func (l *Logger) Debug(v ...interface{}) {
	l.debug.Println(v...)
}
func (l *Logger) Info(v ...interface{}) {
	l.info.Println(v...)
}
func (l *Logger) Warning(v ...interface{}) {
	l.warning.Println(v...)
}
func (l *Logger) Error(v ...interface{}) {
	l.err.Println(v...)
}

// Create Formatted Logs
func (l *Logger) Debugf(format string, v ...interface{}) {
	l.debug.Printf(format, v...)
}
func (l *Logger) Infof(format string, v ...interface{}) {
	l.info.Printf(format, v...)
}
func (l *Logger) Warningf(format string, v ...interface{}) {
	l.warning.Printf(format, v...)
}
func (l *Logger) Errorf(format string, v ...interface{}) {
	l.err.Printf(format, v...)
}
