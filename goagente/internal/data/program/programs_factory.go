// Package programs fornece funcionalidades para coletar informações sobre os programas instalados no sistema.
// Este arquivo contém a definição da estrutura Program, a interface ProgramRetriever,
// e uma fábrica para criar retrievers específicos por sistema operacional.
package programs

import (
	"fmt"
	"runtime"
)

// Program representa informações sobre um programa instalado no sistema.
//
// Campos:
// - Name: Nome do programa.
// - Version: Versão do programa.
type Program struct {
	Name    string `json:"name"`    // Nome do programa
	Version string `json:"version"` // Versão do programa
}

// ProgramRetriever define uma interface para coletar informações sobre programas instalados.
// Implementações específicas para diferentes sistemas operacionais devem implementar esta interface.
type ProgramRetriever interface {
	// GetInstalledPrograms coleta informações sobre os programas instalados no sistema.
	//
	// Retorna:
	// - Um slice de Program contendo as informações dos programas instalados.
	// - Um erro, caso a coleta falhe.
	GetInstalledPrograms() ([]Program, error)
}

// NewProgramRetriever cria e retorna a implementação apropriada de ProgramRetriever
// com base no sistema operacional.
//
// Retorna:
// - Uma implementação de ProgramRetriever compatível com o sistema operacional.
// - Um erro, caso o sistema operacional não seja suportado.
func NewProgramRetriever() (ProgramRetriever, error) {
	so := runtime.GOOS

	switch so {
	case "windows":
		return WindowsProgramsRetriever{}, nil // Implementação para Windows
	case "linux":
		return LinuxProgramsRetriever{}, nil // Implementação para Linux
	default:
		return nil, fmt.Errorf("unsupported operating system: %s", so)
	}
}
