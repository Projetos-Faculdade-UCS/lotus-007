// Package system fornece funcionalidades para coletar informações do sistema.
// Este arquivo contém a implementação de WindowsUserRetriever para coletar o nome do usuário atual no sistema operacional Windows.
package system

import (
	"fmt"
	"goagente/internal/logging"
	"os/exec"
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
	user, err := getInteractiveUser()
	if err != nil {
		logging.Error(fmt.Errorf("erro ao obter o usuário interativo: %w", err))

		return "", err
	} else {
		return user, nil
	}
}

func getInteractiveUser() (string, error) {
	cmd := exec.Command("powershell", "-Command", `(Get-WmiObject -Class Win32_ComputerSystem).UserName`)
	output, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("erro ao executar o comando PowerShell: %w", err)
	}

	user := strings.TrimSpace(string(output))
	if user == "" {
		return "", fmt.Errorf("nenhum usuário interativo encontrado")
	}

	return user, nil
}
