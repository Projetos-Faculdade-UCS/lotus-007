// Package system fornece funcionalidades para coletar informações do sistema.
// Este arquivo contém a definição da interface HostnameRetriever e uma fábrica para criar retrievers específicos por sistema operacional.
package system

import (
	"fmt"
	"runtime"
)

// HostnameRetriever define uma interface para coletar o hostname do sistema.
// Implementações específicas para diferentes sistemas operacionais devem implementar esta interface.
type HostnameRetriever interface {
	// GetHostname coleta o hostname do sistema.
	//
	// Retorna:
	// - Uma string representando o hostname do sistema.
	// - Um erro, caso a coleta falhe.
	GetHostname() (string, error)
}

// NewHostnameRetriever cria e retorna a implementação apropriada de HostnameRetriever
// com base no sistema operacional.
//
// Retorna:
// - Uma implementação de HostnameRetriever compatível com o sistema operacional.
// - Um erro, caso o sistema operacional não seja suportado.
func NewHostnameRetriever() (HostnameRetriever, error) {
	so := runtime.GOOS

	switch so {
	case "windows":
		return WindowsHostnameRetriever{}, nil // Implementação para Windows
	case "linux":
		return LinuxHostnameRetriever{}, nil // Implementação para Linux
	default:
		return nil, fmt.Errorf("sistema operacional não suportado: %s", so)
	}
}
