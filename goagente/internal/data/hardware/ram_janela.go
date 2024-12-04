// Package hardware fornece funcionalidades para coletar informações de hardware do sistema.
// Este arquivo contém a implementação de WindowsRAMRetriever para coletar informações da memória RAM no sistema operacional Windows.
package hardware

import (
	"bytes"
	"encoding/json"
	"fmt"
	"goagente/internal/logging"
	"goagente/pkg/utils"
	"os/exec"
)

// WindowsRAMRetriever é a implementação de RAMRetriever para o sistema operacional Windows.
// Ele utiliza comandos PowerShell para coletar informações sobre a memória RAM.
type WindowsRAMRetriever struct{}

// GetRAMInfo coleta informações sobre a memória RAM no Windows.
// Ele utiliza o PowerShell para obter os dados e os desserializa em uma estrutura RAM.
// As capacidades das memórias são convertidas de bytes para gigabytes.
//
// Retorna:
// - Um slice de RAM contendo as informações das memórias instaladas.
// - Um erro, caso ocorra algum problema durante a execução do comando ou a desserialização dos dados.
func (r WindowsRAMRetriever) GetRAMInfo() ([]RAM, error) {
	cmd := r.powerShellGetRamInfo()

	// Executa o comando PowerShell para obter informações da memória RAM
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		newErr := fmt.Errorf("erro ao executar o comando PowerShell: %v", err)
		logging.Error(newErr)
		return nil, newErr
	}

	// Desserializa a saída JSON em objetos RAM
	ramList, err := r.deserializeRAMInfo(out.Bytes())
	if err != nil {
		logging.Error(err)
		return nil, err
	}

	// Converte a capacidade de cada RAM de bytes para gigabytes
	for i := range ramList {
		ramList[i].Capacity = utils.BytesToGigabytes(uint64(ramList[i].Capacity))
	}

	return ramList, nil
}

// powerShellGetRamInfo cria um comando PowerShell para coletar informações sobre a memória RAM.
// O comando retorna os dados em formato JSON, incluindo Manufacturer, Capacity e FormFactor.
//
// Retorna:
// - Um comando configurado para execução.
func (WindowsRAMRetriever) powerShellGetRamInfo() *exec.Cmd {
	return exec.Command("powershell", "-Command", "Get-WmiObject -Class Win32_PhysicalMemory | Select-Object -Property Manufacturer, Capacity, FormFactor | ConvertTo-Json")
}

// deserializeRAMInfo desserializa os dados JSON em informações da memória RAM.
//
// Funcionalidade:
// - Primeiro, tenta desserializar como um único objeto RAM.
// - Se falhar, tenta desserializar como um array de objetos RAM.
//
// Parâmetros:
// - data: Dados JSON a serem desserializados.
//
// Retorna:
// - Um slice de RAM contendo as informações das memórias.
// - Um erro, caso a desserialização falhe.
func (WindowsRAMRetriever) deserializeRAMInfo(data []byte) ([]RAM, error) {
	// Tenta desserializar o JSON como um único objeto
	var singleRAM RAM
	err := json.Unmarshal(data, &singleRAM)
	if err == nil {
		// Sucesso, retorna como uma lista de um único item
		return []RAM{singleRAM}, nil
	}

	// Se falhar, tenta deserializar como um array de objetos
	var ramList []RAM
	err = json.Unmarshal(data, &ramList)
	if err != nil {
		return nil, fmt.Errorf("erro ao deserializar JSON em RAMInfo: %v", err)
	}

	return ramList, nil
}
