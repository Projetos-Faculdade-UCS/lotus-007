package hardware

import (
	"fmt"
	"runtime"
)

// RAM represents RAM information
type RAM struct {
	Manufacturer string  `json:"Manufacturer"`
	Capacity     float64 `json:"Capacity"` // Capacity in GB
	FormFactor   int     `json:"FormFactor"`
}

// RAMRetriever defines the contract for retrieving RAM information
type RAMRetriever interface {
	GetRAMInfo() ([]RAM, error)
}

// LinuxRAMRetriever is the Linux-specific implementation of RAMRetriever

// NewRAMRetriever returns the correct RAMRetriever implementation based on the OS
func NewRAMRetriever() (RAMRetriever, error) {
	so := runtime.GOOS

	switch so {
	case "windows":
		return WindowsRAMRetriever{}, nil
	case "linux":
		return LinuxRAMRetriever{}, nil
	default:
		return nil, fmt.Errorf("unsupported operating system: %s", so)
	}
}
