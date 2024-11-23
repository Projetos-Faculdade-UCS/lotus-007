// Package logging fornece uma interface para inicializar e gerenciar sistemas de logging.
// Este arquivo contém a implementação específica para o sistema operacional Windows.
package logging

import (
	"log"
	"os"
)

// WindowsLoggerFactory é uma estrutura responsável por inicializar e gerenciar os loggers no sistema operacional Windows.
type WindowsLoggerFactory struct{}

// InitLogger inicializa os loggers para o sistema Windows.
// Ele cria ou abre os arquivos `info.log` e `error.log` para registrar mensagens informativas e de erro.
//
// Retorna:
// - Um erro, se houver falha na criação ou abertura dos arquivos de log.
func (w *WindowsLoggerFactory) InitLogger() error {
	var err error

	// Cria ou abre o arquivo de log para mensagens informativas
	infoLogFile, err = os.OpenFile("info.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return err
	}

	// Cria ou abre o arquivo de log para mensagens de erro
	errorLogFile, err = os.OpenFile("error.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return err
	}

	// Configura os loggers para mensagens informativas e de erro
	infoLogger = log.New(infoLogFile, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	errorLogger = log.New(errorLogFile, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)

	return nil
}

// CloseLogger fecha os arquivos de log criados durante a inicialização.
// Ele garante que os recursos de arquivo sejam liberados corretamente.
func (w *WindowsLoggerFactory) CloseLogger() {
	if infoLogFile != nil {
		infoLogFile.Close()
	}
	if errorLogFile != nil {
		errorLogFile.Close()
	}
}
