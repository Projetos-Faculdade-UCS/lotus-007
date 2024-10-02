package system

import (
	"fmt"
	"runtime"
)

type UserRetriever interface {
	GetCurrentUser() (string, error)
}

func NewUserRetriever() (UserRetriever, error) {
	so := runtime.GOOS

	switch so {
	case "windows":
		return WindowsUserRetriever{}, nil
	case "linux":
		return LinuxUserRetriever{}, nil
	default:
		return nil, fmt.Errorf("sistema operacional não suportado: %s", so)
	}
}
