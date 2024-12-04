// Package system fornece funcionalidades para coletar informações do sistema.
// Este arquivo contém a implementação de WindowsHostnameRetriever para coletar o hostname no sistema operacional Windows.
package system

import (
	"fmt"
	"goagente/internal/logging"
	"os"
)

// WindowsHostnameRetriever é a implementação de HostnameRetriever para o sistema operacional Windows.
// Ele utiliza a biblioteca padrão do Go para obter o hostname do sistema.
type WindowsHostnameRetriever struct{}

// GetHostname coleta o hostname do sistema Windows.
//
// Retorna:
// - Uma string representando o hostname do sistema.
// - Um erro, caso a coleta falhe.
func (WindowsHostnameRetriever) GetHostname() (string, error) {
	hostname, err := os.Hostname()
	if err != nil {
		newErr := fmt.Errorf("erro ao obter o hostname no Windows: %v", err)
		logging.Error(newErr)
		return "", err
	}
	return hostname, nil
}
