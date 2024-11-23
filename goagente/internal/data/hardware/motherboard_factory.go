// Package hardware fornece funcionalidades para coletar informações de hardware do sistema.
// Este arquivo contém a definição de MotherboardInfo e a interface MotherboardInfoRetriever,
// além da fábrica para criar retrievers específicos por sistema operacional.
package hardware

import (
	"fmt"
	"runtime"
)

// MotherboardInfo representa informações sobre a placa-mãe do sistema.
//
// Campos:
// - Manufacturer: Fabricante da placa-mãe.
// - Product: Modelo ou produto da placa-mãe.
type MotherboardInfo struct {
	Manufacturer string `json:"Manufacturer"` // Fabricante da placa-mãe
	Product      string `json:"Product"`      // Modelo ou produto da placa-mãe
}

// MotherboardInfoRetriever define uma interface para coletar informações da placa-mãe.
// Implementações específicas para diferentes sistemas operacionais devem implementar esta interface.
type MotherboardInfoRetriever interface {
	// GetMotherboardInfo coleta informações sobre a placa-mãe do sistema.
	//
	// Retorna:
	// - Um objeto MotherboardInfo contendo as informações da placa-mãe.
	// - Um erro, caso a coleta falhe.
	GetMotherboardInfo() (MotherboardInfo, error)
}

// NewMotherboardRetriever cria e retorna a implementação apropriada de MotherboardInfoRetriever
// com base no sistema operacional.
//
// Retorna:
// - Uma implementação de MotherboardInfoRetriever compatível com o sistema operacional.
// - Um erro, caso o sistema operacional não seja suportado.
func NewMotherboardRetriever() (MotherboardInfoRetriever, error) {
	so := runtime.GOOS

	switch so {
	case "windows":
		return WindowsMotherboardRetriever{}, nil // Implementação para Windows
	case "linux":
		return LinuxMotherboardRetriever{}, nil // Implementação para Linux
	default:
		return nil, fmt.Errorf("unsupported operating system: %s", so)
	}
}
