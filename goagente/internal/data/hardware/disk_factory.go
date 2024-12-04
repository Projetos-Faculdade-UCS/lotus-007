// Package hardware fornece funcionalidades para coletar informações de hardware do sistema.
// Este arquivo contém a definição de DiskInfo e a interface DiskInfoRetriever, além da fábrica para criar retrievers específicos por sistema operacional.
package hardware

import (
	"fmt"
	"runtime"
)

// DiskInfo representa informações sobre um disco físico no sistema.
//
// Campos:
// - DeviceID: Identificação do dispositivo.
// - Model: Modelo do disco.
// - Size: Tamanho do disco em bytes (uint64).
type DiskInfo struct {
	DeviceID string `json:"DeviceID"` // Identificação do dispositivo
	Model    string `json:"Model"`    // Modelo do disco
	Size     uint64 `json:"Size"`     // Tamanho do disco em bytes
}

// DiskInfoRetriever define uma interface para coletar informações de discos físicos.
// Implementações específicas para sistemas operacionais devem implementar esta interface.
type DiskInfoRetriever interface {
	// GetDiskInfo coleta informações sobre os discos físicos disponíveis no sistema.
	//
	// Retorna:
	// - Um slice de DiskInfo contendo as informações dos discos.
	// - Um erro, caso a coleta falhe.
	GetDiskInfo() ([]DiskInfo, error)
}

// NewDiskRetriever cria e retorna a implementação apropriada de DiskInfoRetriever com base no sistema operacional.
//
// Retorna:
// - Uma implementação de DiskInfoRetriever compatível com o sistema operacional.
// - Um erro, caso o sistema operacional não seja suportado.
func NewDiskRetriever() (DiskInfoRetriever, error) {
	so := runtime.GOOS

	switch so {
	case "windows":
		return WindowsDiskRetriever{}, nil // Implementação para Windows
	case "linux":
		return LinuxDiskRetriever{}, nil // Implementação para Linux
	default:
		return nil, fmt.Errorf("unsupported operating system: %s", so)
	}
}
