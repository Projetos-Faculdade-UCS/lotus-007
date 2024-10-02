package hardware

import (
	"fmt"
	"runtime"
)

type MotherboardInfo struct {
	Manufacturer string `json:"Manufacturer"`
	Product      string `json:"Product"`
}

type MotherboardInfoRetriever interface {
	GetMotherboardInfo() (MotherboardInfo, error)
}

func NewMotherboardRetriever() (MotherboardInfoRetriever, error) {
	so := runtime.GOOS

	switch so {
	case "windows":
		return WindowsMotherboardRetriever{}, nil
	case "linux":
		return LinuxMotherboardRetriever{}, nil
	default:
		return nil, fmt.Errorf("unsupported operating system: %s", so)
	}
}
