package system

import (
	"fmt"
	"goagente/internal/logging"
	"os"
)

type LinuxHostnameRetriever struct{}

func (LinuxHostnameRetriever) GetHostname() (string, error) {
	hostname, err := os.Hostname()
	if err != nil {
		newErr := fmt.Errorf("erro ao obter o hostname no Linux: %v", err)
		logging.Error(newErr)
		return "", err
	}
	return hostname, nil
}
