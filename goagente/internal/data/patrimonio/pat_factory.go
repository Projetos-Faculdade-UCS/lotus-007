package patrimonio

import (
	"fmt"
	"runtime"
)

type PatRetriever interface {
	GetCurrentPat() (string, error)
}

func NewPatRetriever() (PatRetriever, error) {
	so := runtime.GOOS

	switch so {
	case "windows":
		return WindowsPatRetriever{}, nil
	case "linux":
		return LinuxPatRetriever{}, nil
	default:
		return nil, fmt.Errorf("sistema operacional não suportado: %s", so)
	}
}
