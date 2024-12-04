// Package system fornece funcionalidades para coletar informações do sistema.
// Este arquivo contém a definição da interface OSRetriever e uma fábrica para criar retrievers específicos por sistema operacional.
package system

import (
	"fmt"
	"runtime"
)

// OSRetriever define uma interface para coletar informações sobre o sistema operacional atual.
// Implementações específicas para diferentes sistemas operacionais devem implementar esta interface.
type OSRetriever interface {
	// GetCurrentOS coleta informações sobre o sistema operacional atual.
	//
	// Retorna:
	// - Uma string representando o sistema operacional.
	// - Um erro, caso a coleta falhe.
	GetCurrentOS() (string, error)
}

// NewOSRetriever cria e retorna a implementação apropriada de OSRetriever
// com base no sistema operacional.
//
// Retorna:
// - Uma implementação de OSRetriever compatível com o sistema operacional.
// - Um erro, caso o sistema operacional não seja suportado.
func NewOSRetriever() (OSRetriever, error) {
	so := runtime.GOOS

	switch so {
	case "windows":
		return WindowsOSRetriever{}, nil // Implementação para Windows
	case "linux":
		return LinuxOSRetriever{}, nil // Implementação para Linux
	default:
		return nil, fmt.Errorf("sistema operacional não suportado: %s", so)
	}
}
