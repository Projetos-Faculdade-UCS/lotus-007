// Package system fornece funcionalidades para coletar informações do sistema.
// Este arquivo contém a implementação de WindowsUserRetriever para coletar o nome do usuário atual no sistema operacional Windows.
package system

import (
	"fmt"
	"goagente/internal/logging"
	"os/user"
	"strings"
)

// WindowsUserRetriever é a implementação de UserRetriever para o sistema operacional Windows.
// Ele utiliza a biblioteca padrão do Go para obter o nome do usuário atual.
type WindowsUserRetriever struct{}

// GetCurrentUser coleta o nome do usuário atual do sistema Windows.
//
// Retorna:
// - Uma string representando o nome do usuário atual.
// - Um erro, caso a coleta falhe.
func (WindowsUserRetriever) GetCurrentUser() (string, error) {
	currentUser, err := user.Current()
	if err != nil {
		newErr := fmt.Errorf("erro ao obter o usuário atual no Windows: %v", err)
		logging.Error(newErr)
		return "", err
	}

	username := currentUser.Username
	// Verifica se o nome do usuário contém um domínio (formato DOMAIN\Username)
	if strings.Contains(username, "\\") {
		parts := strings.Split(username, "\\")
		if len(parts) > 0 {
			// Retorna apenas o nome de usuário, ignorando o domínio
			return parts[len(parts)-1], nil
		}
	}

	return username, nil
}
