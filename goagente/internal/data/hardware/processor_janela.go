// Package hardware fornece funcionalidades para coletar informações de hardware do sistema.
// Este arquivo contém a implementação do WindowsProcessorRetriever para coletar dados do processador no sistema operacional Windows.
package hardware

import (
	"bytes"
	"encoding/json"
	"fmt"
	"goagente/internal/logging"
	"os/exec"
)

// WindowsProcessorRetriever é a implementação de ProcessorInfoRetriever para o sistema operacional Windows.
// Ele utiliza comandos PowerShell para coletar informações sobre os processadores.
type WindowsProcessorRetriever struct{}

// GetProcessorInfo coleta informações sobre os processadores no Windows.
// Ele utiliza o PowerShell para obter os dados e os desserializa em uma estrutura ProcessorInfo.
//
// Retorna:
// - Um slice de ProcessorInfo contendo as informações dos processadores.
// - Um erro, caso ocorra algum problema durante a execução do comando ou a desserialização dos dados.
func (p WindowsProcessorRetriever) GetProcessorInfo() ([]ProcessorInfo, error) {
	cmd := p.powerShellGetProcessorInfo()

	// Executa o comando PowerShell para obter informações dos processadores
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		newErr := fmt.Errorf("erro ao executar o comando PowerShell em GetProcessorInfo: %v", err)
		logging.Error(newErr)
		return nil, err
	}

	// Desserializa a saída JSON em objetos ProcessorInfo
	processors, err := p.deserializeProcessorInfo(out.Bytes())
	if err != nil {
		logging.Error(err)
		return nil, err
	}

	return processors, nil
}

// powerShellGetProcessorInfo cria um comando PowerShell para coletar informações sobre os processadores.
// O comando retorna os dados em formato JSON, incluindo Name, NumberOfCores e MaxClockSpeed.
//
// Retorna:
// - Um comando configurado para execução.
func (WindowsProcessorRetriever) powerShellGetProcessorInfo() *exec.Cmd {
	// Comando PowerShell para obter informações do processador em formato JSON
	cmd := exec.Command("powershell", "-Command", "Get-WmiObject -Class Win32_Processor | Select-Object -Property Name, NumberOfCores, MaxClockSpeed | ConvertTo-Json")
	return cmd
}

// deserializeProcessorInfo desserializa os dados JSON em um ou mais objetos ProcessorInfo.
//
// Funcionalidade:
// - Primeiro, tenta desserializar como um único objeto ProcessorInfo.
// - Se falhar, tenta desserializar como um array de objetos ProcessorInfo.
//
// Parâmetros:
// - data: Dados JSON a serem desserializados.
//
// Retorna:
// - Um slice de ProcessorInfo contendo as informações dos processadores.
// - Um erro, caso a desserialização falhe.
func (p WindowsProcessorRetriever) deserializeProcessorInfo(data []byte) ([]ProcessorInfo, error) {
	// Tenta desserializar o JSON como um único objeto
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
