package utils

import (
	"os"
	"io"
	"log"
	"github.com/shappy0/saasc/internal/config"
)

const (
	DefaultDirMod os.FileMode = 0755
	DefaultFileMod os.FileMode = 0600
	Info = "info"
	Warning = "warning"
	Error = "error"
)

type Logger struct {
	FilePath	string
	File		*os.File
	LInfo 		*log.Logger
	LWarning 	*log.Logger
	LError 		*log.Logger
}

func NewLogger(c *config.Conf) (*Logger, error) {
	filePath := c.LogDirPath
	lr := NewLogRotator(filePath, c.LogDirPath)
	file := lr.File
	multiWriter := io.MultiWriter(file) //file, os.Stdout
	multiWriterErr := io.MultiWriter(file) //file, os.Stderr
	l := &Logger{
		FilePath:	filePath,
		File:		file,
		LInfo:		log.New(multiWriter, "INFO: ", log.Ldate|log.Ltime),
		LWarning:	log.New(multiWriter, "WARNING: ", log.Ldate|log.Ltime),
		LError:		log.New(multiWriterErr, "ERROR: ", log.Ldate|log.Ltime),
	}
	l.LInfo.SetOutput(lr)
	return l, nil
}

func createLogFile(filePath string) (*os.File, error) {
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, DefaultFileMod)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func (l *Logger) Log(kind, message string) {
	switch kind {
	case Info:
		l.LInfo.Println(message)
	case Warning:
		l.LWarning.Println(message)
	case Error:
		l.LError.Println(message)
	default:
		log.Println("Unsupported log type %s", kind)
	}
}

func (l *Logger) Info(message string) {
	l.LInfo.Println(message)
}

func (l *Logger) Infof(format string, v ...any) {
	l.LInfo.Printf(format, v...)
}

func (l *Logger) Warn(message string) {
	l.LWarning.Println(message)
}

func (l *Logger) Warnf(format string, v ...any) {
	l.LWarning.Printf(format, v...)
}

func (l *Logger) Warning(message string) {
	l.LWarning.Println(message)
}

func (l *Logger) Warningf(format string, v ...any) {
	l.LWarning.Printf(format, v...)
}

func (l *Logger) Error(message string) {
	l.LError.Println(message)
}

func (l *Logger) Errorf(format string, v ...any) {
	l.LError.Printf(format, v...)
}