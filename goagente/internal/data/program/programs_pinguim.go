package programs

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

type LinuxProgramsRetriever struct{}

// GetInstalledPrograms obtém a lista de programas instalados no Linux usando dpkg-query
func (p LinuxProgramsRetriever) GetInstalledPrograms() ([]Program, error) {
	var programs []Program

	// Executa o comando dpkg-query para listar pacotes instalados
	cmd := exec.Command("dpkg-query", "-W", "-f=${Package} ${Version}\n")

	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return nil, fmt.Errorf("erro ao executar dpkg-query: %v", err)
	}

	// Processa a saída do comando dpkg-query
	lines := strings.Split(out.String(), "\n")
	for _, line := range lines {
		if line == "" {
			continue
		}

		// Divide a linha em nome e versão do pacote
		parts := strings.SplitN(line, " ", 2)
		if len(parts) == 2 {
			program := Program{
				Name:    parts[0],
				Version: parts[1],
			}
			programs = append(programs, program)
		}
	}

	return programs, nil
}
