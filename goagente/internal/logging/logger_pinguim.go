package logging

import (
	"log"
	"os"
)

var (
	infoLogger   *log.Logger
	errorLogger  *log.Logger
	infoLogFile  *os.File
	errorLogFile *os.File
)

type LinuxLoggerFactory struct{}

// InitLogger inicializa os loggers no Linux
func (l *LinuxLoggerFactory) InitLogger() error {
	var err error

	infoLogFile, err = os.OpenFile("info.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return err
	}

	errorLogFile, err = os.OpenFile("error.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return err
	}

	infoLogger = log.New(infoLogFile, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	errorLogger = log.New(errorLogFile, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)

	return nil
}

// CloseLogger fecha os arquivos de log no Linux
func (l *LinuxLoggerFactory) CloseLogger() {
	if infoLogFile != nil {
		infoLogFile.Close()
	}
	if errorLogFile != nil {
		errorLogFile.Close()
	}
}
