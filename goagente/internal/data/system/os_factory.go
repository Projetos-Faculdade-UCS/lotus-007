package system

import (
	"fmt"
	"runtime"
)

type OSRetriever interface {
	GetCurrentOS() (string, error)
}

func NewOSRetriever() (OSRetriever, error) {
	so := runtime.GOOS

	switch so {
	case "windows":
		return WindowsOSRetriever{}, nil
	case "linux":
		return LinuxOSRetriever{}, nil
	default:
		return nil, fmt.Errorf("sistema operacional não suportado: %s", so)
	}
}
