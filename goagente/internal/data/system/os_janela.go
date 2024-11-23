// Package system fornece funcionalidades para coletar informações do sistema.
// Este arquivo contém a implementação de WindowsOSRetriever para coletar informações sobre o sistema operacional Windows.
package system

import (
	"fmt"
	"os/exec"
	"strings"
)

// WindowsOSRetriever é a implementação de OSRetriever para o sistema operacional Windows.
// Ele utiliza comandos do sistema para coletar informações sobre a versão do Windows.
type WindowsOSRetriever struct{}

// GetCurrentOS coleta informações sobre o sistema operacional Windows.
//
// Retorna:
// - Uma string representando o nome e a versão do sistema operacional.
// - Um erro, caso a coleta falhe.
func (WindowsOSRetriever) GetCurrentOS() (string, error) {
	// Comando para obter o nome do sistema operacional
	cmd := exec.Command("cmd", "/C", "wmic os get Caption")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}

	// Converte a saída para string, remove espaços extras e separa em linhas
	lines := strings.Split(strings.TrimSpace(string(output)), "\n")

	// Verifica se há uma linha com o Caption e retorna a segunda linha, que contém o valor
	if len(lines) >= 2 {
		return strings.TrimSpace(lines[1]), nil
	}

	return "", fmt.Errorf("erro em pegar a versão do Windows")
}
