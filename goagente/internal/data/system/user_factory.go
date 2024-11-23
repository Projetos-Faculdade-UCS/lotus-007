// Package system fornece funcionalidades para coletar informações do sistema.
// Este arquivo contém a definição da interface UserRetriever e uma fábrica para criar retrievers específicos por sistema operacional.
package system

import (
	"fmt"
	"runtime"
)

// UserRetriever define uma interface para coletar o usuário atual do sistema.
// Implementações específicas para diferentes sistemas operacionais devem implementar esta interface.
type UserRetriever interface {
	// GetCurrentUser coleta o nome do usuário atual do sistema.
	//
	// Retorna:
	// - Uma string representando o nome do usuário atual.
	// - Um erro, caso a coleta falhe.
	GetCurrentUser() (string, error)
}

// NewUserRetriever cria e retorna a implementação apropriada de UserRetriever
// com base no sistema operacional.
//
// Retorna:
// - Uma implementação de UserRetriever compatível com o sistema operacional.
// - Um erro, caso o sistema operacional não seja suportado.
func NewUserRetriever() (UserRetriever, error) {
	so := runtime.GOOS

	switch so {
	case "windows":
		return WindowsUserRetriever{}, nil // Implementação para Windows
	case "linux":
		return LinuxUserRetriever{}, nil // Implementação para Linux
	default:
		return nil, fmt.Errorf("sistema operacional não suportado: %s", so)
	}
}
