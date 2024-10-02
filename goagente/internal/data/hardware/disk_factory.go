package hardware

import (
	"fmt"
	"runtime"
)

type DiskInfo struct {
	DeviceID string `json:"DeviceID"`
	Model    string `json:"Model"`
	Size     uint64 `json:"Size"` // Mantendo como uint64
}
type DiskInfoRetriever interface {
	GetDiskInfo() ([]DiskInfo, error)
}

func NewDiskRetriever() (DiskInfoRetriever, error) {
	so := runtime.GOOS

	switch so {
	case "windows":
		return WindowsDiskRetriever{}, nil
	case "linux":
		return LinuxDiskRetriever{}, nil
	default:
		return nil, fmt.Errorf("unsupported operating system: %s", so)
	}
}
