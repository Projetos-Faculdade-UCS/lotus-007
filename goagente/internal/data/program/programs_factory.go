package programs

import (
	"fmt"
	"runtime"
)

type Program struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

type ProgramRetriever interface {
	GetInstalledPrograms() ([]Program, error)
}

func NewProgramRetriever() (ProgramRetriever, error) {
	so := runtime.GOOS

	switch so {
	case "windows":
		return WindowsProgramsRetriever{}, nil
	case "linux":
		return LinuxProgramsRetriever{}, nil
	default:
		return nil, fmt.Errorf("unsupported operating system: %s", so)
	}
}
