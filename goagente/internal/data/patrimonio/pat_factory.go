// Package patrimonio fornece funcionalidades para coletar informações sobre o patrimônio do sistema.
// Este arquivo contém a definição da interface PatRetriever e uma fábrica para criar retrievers específicos por sistema operacional.
package patrimonio

import (
	"fmt"
	"runtime"
)

// PatRetriever define uma interface para coletar o patrimônio do sistema.
// Implementações específicas para diferentes sistemas operacionais devem implementar esta interface.
type PatRetriever interface {
	// GetCurrentPat coleta o patrimônio atual do sistema.
	//
	// Retorna:
	// - Uma string representando o patrimônio do sistema.
	// - Um erro, caso a coleta falhe.
	GetCurrentPat() (string, error)
}

// NewPatRetriever cria e retorna a implementação apropriada de PatRetriever
// com base no sistema operacional.
//
// Retorna:
// - Uma implementação de PatRetriever compatível com o sistema operacional.
// - Um erro, caso o sistema operacional não seja suportado.
func NewPatRetriever() (PatRetriever, error) {
	so := runtime.GOOS

	switch so {
	case "windows":
		return WindowsPatRetriever{}, nil // Implementação para Windows
	case "linux":
		return LinuxPatRetriever{}, nil // Implementação para Linux
	default:
		return nil, fmt.Errorf("sistema operacional não suportado: %s", so)
	}
}
