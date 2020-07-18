package log

import (
	"fmt"
	"io"
	"log"
)

type LogLevel int

const (
	LOG_DEBUG LogLevel = iota
	LOG_INFO
	LOG_WARNING
	LOG_ERR
	LOG_OFF
	LOG_UNKNOWN
)

const (
	DEFAULT_LOG_PREFIX = "[goant]"
	DEFAULT_LOG_FLAG   = log.Ldate | log.Lmicroseconds
	DEFAULT_LOG_LEVEL  = LOG_DEBUG
)

type Logger interface {
	Info(v ...interface{})

	Infof(format string, v ...interface{})

	Warn(v ...interface{})

	Warnf(format string, v ...interface{})

	Debug(v ...interface{})

	Debugf(format string, v ...interface{})

	Err(v ...interface{})

	Errf(format string, v ...interface{})

	Level() LogLevel

	SetLevel(l LogLevel)
}

type SimpleLogger struct {
	logger *log.Logger
	level  LogLevel
}

func (s *SimpleLogger) Info(v ...interface{}) {

	if s.level <= LOG_INFO {
		s.logger.Output(2, fmt.Sprintln(v...))
	}

	return
}

func (s *SimpleLogger) Infof(format string, v ...interface{}) {

	if s.level <= LOG_INFO {
		s.logger.Output(2, fmt.Sprintf(format, v...))
	}

	return
}

func (s *SimpleLogger) Warn(v ...interface{}) {

	if s.level <= LOG_WARNING {
		s.logger.Output(2, fmt.Sprintln(v...))
	}

	return
}

func (s *SimpleLogger) Warnf(format string, v ...interface{}) {

	if s.level <= LOG_WARNING {
		s.logger.Output(2, fmt.Sprintf(format, v...))
	}

	return
}

func (s *SimpleLogger) Debug(v ...interface{}) {

	if s.level <= LOG_DEBUG {
		s.logger.Output(2, fmt.Sprintln(v...))
	}

	return
}

func (s *SimpleLogger) Debugf(format string, v ...interface{}) {

	if s.level <= LOG_DEBUG {
		s.logger.Output(2, fmt.Sprintf(format, v...))
	}

	return
}

func (s *SimpleLogger) Err(v ...interface{}) {

	if s.level <= LOG_ERR {
		s.logger.Output(2, fmt.Sprintln(v...))
	}

	return
}

func (s *SimpleLogger) Errf(format string, v ...interface{}) {

	if s.level <= LOG_ERR {
		s.logger.Output(2, fmt.Sprintf(format, v...))
	}

	return
}

func (s *SimpleLogger) Level() LogLevel {

	return s.level
}

func (s *SimpleLogger) SetLevel(l LogLevel) {

	s.level = l
	return

}

func New(writer io.Writer) *SimpleLogger {

	return &SimpleLogger{

		logger: log.New(writer, DEFAULT_LOG_PREFIX, DEFAULT_LOG_FLAG),
		level:  DEFAULT_LOG_LEVEL,
	}
}

func New2(writer io.Writer, prefix string, flag int) *SimpleLogger {

	return &SimpleLogger{

		logger: log.New(writer, prefix, flag),
		level:  LOG_DEBUG,
	}
}
