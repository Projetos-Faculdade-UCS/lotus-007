package system

import (
	"fmt"
	"goagente/internal/logging"
	"os"
)

type WindowsHostnameRetriever struct{}

func (WindowsHostnameRetriever) GetHostname() (string, error) {
	hostname, err := os.Hostname()
	if err != nil {
		newErr := fmt.Errorf("erro ao obter o hostname no Windows: %v", err)
		logging.Error(newErr)
		return "", err
	}
	return hostname, nil
}
