// Package hardware fornece funcionalidades para coletar informações de hardware do sistema.
// Este arquivo contém a definição de ProcessorInfo e a interface ProcessorInfoRetriever,
// além da fábrica para criar retrievers específicos por sistema operacional.
package hardware

import (
	"fmt"
	"runtime"
)

// ProcessorInfo representa informações sobre os processadores do sistema.
//
// Campos:
// - Name: Nome do processador.
// - NumberOfCores: Número de núcleos disponíveis no processador.
// - MaxClockSpeed: Velocidade máxima do clock do processador, em MHz.
type ProcessorInfo struct {
	Name          string `json:"Name"`          // Nome do processador
	NumberOfCores int    `json:"NumberOfCores"` // Número de núcleos
	MaxClockSpeed int    `json:"MaxClockSpeed"` // Velocidade máxima do clock em MHz
}

// ProcessorInfoRetriever define uma interface para coletar informações de processadores.
// Implementações específicas para diferentes sistemas operacionais devem implementar esta interface.
type ProcessorInfoRetriever interface {
	// GetProcessorInfo coleta informações sobre os processadores do sistema.
	//
	// Retorna:
	// - Um slice de ProcessorInfo contendo as informações dos processadores.
	// - Um erro, caso a coleta falhe.
	GetProcessorInfo() ([]ProcessorInfo, error)
}

// NewProcessorRetriever cria e retorna a implementação apropriada de ProcessorInfoRetriever
// com base no sistema operacional.
//
// Retorna:
// - Uma implementação de ProcessorInfoRetriever compatível com o sistema operacional.
// - Um erro, caso o sistema operacional não seja suportado.
func NewProcessorRetriever() (ProcessorInfoRetriever, error) {
	so := runtime.GOOS

	switch so {
	case "windows":
		return WindowsProcessorRetriever{}, nil // Implementação para Windows
	case "linux":
		return LinuxProcessorRetriever{}, nil // Implementação para Linux
	default:
		return nil, fmt.Errorf("unsupported operating system: %s", so)
	}
}
