package system

import (
	"fmt"
	"runtime"
)

type HostnameRetriever interface {
	GetHostname() (string, error)
}

func NewHostnameRetriever() (HostnameRetriever, error) {
	so := runtime.GOOS

	switch so {
	case "windows":
		return WindowsHostnameRetriever{}, nil
	case "linux":
		return LinuxHostnameRetriever{}, nil
	default:
		return nil, fmt.Errorf("sistema operacional não suportado: %s", so)
	}
}
