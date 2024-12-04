// Package hardware fornece funcionalidades para coletar informações de hardware do sistema.
// Este arquivo contém a definição de RAM, a interface RAMRetriever, e uma fábrica para criar
// retrievers específicos por sistema operacional.
package hardware

import (
	"fmt"
	"runtime"
)

// RAM representa informações sobre a memória RAM do sistema.
//
// Campos:
// - Manufacturer: Fabricante da memória RAM.
// - Capacity: Capacidade da memória em gigabytes (GB).
// - FormFactor: Formato físico da memória RAM.
type RAM struct {
	Manufacturer string  `json:"Manufacturer"` // Fabricante da memória RAM
	Capacity     float64 `json:"Capacity"`     // Capacidade da memória em GB
	FormFactor   int     `json:"FormFactor"`   // Formato físico da memória RAM
}

// RAMRetriever define uma interface para coletar informações da memória RAM.
// Implementações específicas para diferentes sistemas operacionais devem implementar esta interface.
type RAMRetriever interface {
	// GetRAMInfo coleta informações sobre a memória RAM do sistema.
	//
	// Retorna:
	// - Um slice de RAM contendo as informações das memórias instaladas.
	// - Um erro, caso a coleta falhe.
	GetRAMInfo() ([]RAM, error)
}

// NewRAMRetriever cria e retorna a implementação apropriada de RAMRetriever
// com base no sistema operacional.
//
// Retorna:
// - Uma implementação de RAMRetriever compatível com o sistema operacional.
// - Um erro, caso o sistema operacional não seja suportado.
func NewRAMRetriever() (RAMRetriever, error) {
	so := runtime.GOOS

	switch so {
	case "windows":
		return WindowsRAMRetriever{}, nil // Implementação para Windows
	case "linux":
		return LinuxRAMRetriever{}, nil // Implementação para Linux
	default:
		return nil, fmt.Errorf("unsupported operating system: %s", so)
	}
}
