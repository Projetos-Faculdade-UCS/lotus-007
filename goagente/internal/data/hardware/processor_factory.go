package hardware

import (
	"fmt"
	"runtime"
)

type ProcessorInfo struct {
	Name          string `json:"Name"`
	NumberOfCores int    `json:"NumberOfCores"`
	MaxClockSpeed int    `json:"MaxClockSpeed"`
}

type ProcessorInfoRetriever interface {
	GetProcessorInfo() ([]ProcessorInfo, error)
}

func NewProcessorRetriever() (ProcessorInfoRetriever, error) {
	so := runtime.GOOS

	switch so {
	case "windows":
		return WindowsProcessorRetriever{}, nil
	case "linux":
		return LinuxProcessorRetriever{}, nil
	default:
		return nil, fmt.Errorf("unsupported operating system: %s", so)
	}
}
