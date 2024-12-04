// Package logging fornece uma interface para inicializar e gerenciar sistemas de logging.
// Ele utiliza o padrão Singleton para garantir que apenas uma instância de LoggerFactory seja criada.
package logging

import (
	"fmt"
	"runtime"
	"sync"
)

var (
	once      sync.Once     // Garante a inicialização única do Singleton
	singleton LoggerFactory // Instância única de LoggerFactory
	initError error         // Erro de inicialização, caso ocorra
)

// LoggerFactory define os métodos que um sistema de logging deve implementar.
//
// Métodos:
// - InitLogger(): Inicializa o logger.
// - CloseLogger(): Finaliza e limpa os recursos do logger.
type LoggerFactory interface {
	InitLogger() error
	CloseLogger()
}

// NewLoggerFactory retorna a única instância de LoggerFactory (Singleton).
// A implementação retornada depende do sistema operacional.
//
// Retorna:
// - Uma instância de LoggerFactory, se suportado.
// - Um erro, caso o sistema operacional não seja suportado.
func NewLoggerFactory() (LoggerFactory, error) {
	once.Do(func() {
		so := runtime.GOOS // Identifica o sistema operacional
		switch so {
		case "windows":
			singleton = &WindowsLoggerFactory{}
		case "linux":
			singleton = &LinuxLoggerFactory{}
		default:
			initError = fmt.Errorf("sistema operacional não suportado: %s", so)
		}

		// Inicializa o logger, caso o sistema operacional seja suportado
		if initError == nil {
			initError = singleton.InitLogger()
		}
	})

	return singleton, initError
}

// Info registra uma mensagem informativa no logger.
//
// Parâmetros:
// - message: A mensagem informativa a ser registrada.
func Info(message string) {
	infoLogger.Println(message)
}

// Error registra um erro no logger.
//
// Parâmetros:
// - err: O erro a ser registrado.
func Error(err error) {
	errorLogger.Println(err)
}

// Debug registra uma mensagem de depuração no logger.
//
// Parâmetros:
// - message: A mensagem de depuração a ser registrada.
func Debug(message string) {
	infoLogger.Println("DEBUG: ", message)
}
