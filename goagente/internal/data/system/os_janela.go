package system

import (
	"fmt"
	"os/exec"
	"strings"
)

type WindowsOSRetriever struct{}

func (WindowsOSRetriever) GetCurrentOS() (string, error) {
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
