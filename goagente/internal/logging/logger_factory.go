package logging

import (
	"fmt"
	"runtime"
	"sync"
)

var (
	once      sync.Once
	singleton LoggerFactory
	initError error
)

// LoggerFactory interface define os métodos de inicialização e fechamento
type LoggerFactory interface {
	InitLogger() error
	CloseLogger()
}

// NewLoggerFactory retorna a única instância de LoggerFactory (Singleton)
func NewLoggerFactory() (LoggerFactory, error) {
	once.Do(func() {
		so := runtime.GOOS
		switch so {
		case "windows":
			singleton = &WindowsLoggerFactory{}
		case "linux":
			singleton = &LinuxLoggerFactory{}
		default:
			initError = fmt.Errorf("sistema operacional não suportado: %s", so)
		}

		if initError == nil {
			initError = singleton.InitLogger()
		}
	})

	return singleton, initError
}

func Info(message string) {
	infoLogger.Println(message)
}

func Error(err error) {
	errorLogger.Println(err)
}

func Debug(message string) {
	infoLogger.Println("DEBUG: ", message)
}
