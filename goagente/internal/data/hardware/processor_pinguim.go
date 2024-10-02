package hardware

import (
	"bytes"
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

// LinuxProcessorRetriever é o struct para obter informações do processador no Linux
type LinuxProcessorRetriever struct{}

// GetProcessorInfo obtém informações do processador no Linux
func (l LinuxProcessorRetriever) GetProcessorInfo() ([]ProcessorInfo, error) {
	cmd := exec.Command("lscpu")

	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return nil, fmt.Errorf("erro ao executar o comando lscpu: %v", err)
	}

	// Processa a saída do comando lscpu
	processors, err := l.parseLSCPUOutput(out.String())
	if err != nil {
		return nil, err
	}

	return processors, nil
}

// parseLSCPUOutput processa a saída do comando lscpu e extrai as informações necessárias
func (l LinuxProcessorRetriever) parseLSCPUOutput(data string) ([]ProcessorInfo, error) {
	var processorInfo ProcessorInfo
	lines := strings.Split(data, "\n")

	for _, line := range lines {
		line = strings.TrimSpace(line)

		// Extrai o nome do processador
		if strings.HasPrefix(line, "Model name:") {
			processorInfo.Name = strings.TrimSpace(strings.TrimPrefix(line, "Model name:"))
		}

		// Extrai o número de núcleos
		if strings.HasPrefix(line, "CPU(s):") {
			coreCount, err := strconv.Atoi(strings.TrimSpace(strings.TrimPrefix(line, "CPU(s):")))
			if err != nil {
				return nil, fmt.Errorf("erro ao converter número de núcleos: %v", err)
			}
			processorInfo.NumberOfCores = coreCount
		}

		// Extrai a velocidade máxima do clock
		if strings.HasPrefix(line, "CPU max MHz:") {
			maxClockSpeed, err := strconv.Atoi(strings.TrimSpace(strings.TrimPrefix(line, "CPU max MHz:")))
			if err != nil {
				return nil, fmt.Errorf("erro ao converter velocidade máxima do clock: %v", err)
			}
			processorInfo.MaxClockSpeed = maxClockSpeed
		}
	}

	// Verifica se as informações foram coletadas
	if processorInfo.Name == "" || processorInfo.NumberOfCores == 0 || processorInfo.MaxClockSpeed == 0 {
		return nil, fmt.Errorf("informações do processador incompletas")
	}

	// Retorna as informações como um slice com um único item (normalmente um processador no Linux)
	return []ProcessorInfo{processorInfo}, nil
}
