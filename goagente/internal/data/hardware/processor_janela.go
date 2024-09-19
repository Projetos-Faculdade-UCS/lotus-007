package hardware

import (
	"bytes"
	"encoding/json"
	"fmt"
	"goagente/internal/logging"
	"os/exec"
)

// Estrutura para informações do processador

// ProcessorInfoRetrieverWindows define o contrato para obter informações do processador
type WindowsProcessorRetriever struct{}

// GetProcessorInfo retorna as informações do processador
func (p WindowsProcessorRetriever) GetProcessorInfo() ([]ProcessorInfo, error) {
	cmd := p.powerShellGetProcessorInfo()

	// Executa o comando PowerShell
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		newErr := fmt.Errorf("erro ao executar o comando PowerShell em GetProcessorInfo: %v", err)
		logging.Error(newErr)
		return nil, err
	}

	// Usa o novo método para desserializar o JSON
	processors, err := p.deserializeProcessorInfo(out.Bytes())
	if err != nil {
		logging.Error(err)
		return nil, err
	}

	return processors, nil
}

// PowerShellGetProcessorInfo executa o comando PowerShell para obter informações do processador
func (WindowsProcessorRetriever) powerShellGetProcessorInfo() *exec.Cmd {
	// Comando PowerShell para obter informações do processador em formato JSON
	cmd := exec.Command("powershell", "-Command", "Get-WmiObject -Class Win32_Processor | Select-Object -Property Name, NumberOfCores, MaxClockSpeed | ConvertTo-Json")
	return cmd
}

// deserializeProcessorInfo tenta desserializar o JSON como um único objeto ou uma lista de objetos ProcessorInfo
func (p WindowsProcessorRetriever) deserializeProcessorInfo(data []byte) ([]ProcessorInfo, error) {
	// Tenta deserializar o JSON como um único objeto
	var singleProcessor ProcessorInfo
	err := json.Unmarshal(data, &singleProcessor)
	if err == nil {
		// Se bem-sucedido, retorna como um slice com um único item
		return []ProcessorInfo{singleProcessor}, nil
	}

	// Se falhar, tenta deserializar como uma lista de objetos
	var processors []ProcessorInfo
	err = json.Unmarshal(data, &processors)
	if err != nil {
		return nil, fmt.Errorf("erro ao deserializar JSON em ProcessorInfo: %v", err)
	}

	return processors, nil
}
