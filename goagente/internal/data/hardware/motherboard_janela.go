// Package hardware fornece funcionalidades para coletar informações sobre hardware do sistema.
// Este arquivo contém a implementação do WindowsMotherboardRetriever para coletar dados da placa-mãe no sistema operacional Windows.
package hardware

import (
	"bytes"
	"encoding/json"
	"fmt"
	"goagente/internal/logging"
	"os/exec"
)

// WindowsMotherboardRetriever é a implementação de MotherboardInfoRetriever para o sistema operacional Windows.
// Ele utiliza comandos PowerShell para coletar informações sobre a placa-mãe.
type WindowsMotherboardRetriever struct{}

// GetMotherboardInfo coleta informações sobre a placa-mãe no Windows.
// Ele utiliza o PowerShell para obter os dados e os desserializa em uma estrutura MotherboardInfo.
//
// Retorna:
// - Um objeto MotherboardInfo contendo as informações da placa-mãe.
// - Um erro, caso ocorra algum problema durante a execução do comando ou a desserialização dos dados.
func (r WindowsMotherboardRetriever) GetMotherboardInfo() (MotherboardInfo, error) {
	cmd := r.powerShellGetMotherboardInfo()

	// Executa o comando PowerShell para obter informações da placa-mãe
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		newErr := fmt.Errorf("erro ao executar o comando PowerShell em GetMotherboardInfo: %v", err)
		logging.Error(newErr)
		return MotherboardInfo{}, err
	}

	// Desserializa a saída JSON em um objeto MotherboardInfo
	motherboard, err := r.deserializeMotherboardInfo(out.Bytes())
	if err != nil {
		logging.Error(err)
		return MotherboardInfo{}, err
	}

	return motherboard, nil
}

// powerShellGetMotherboardInfo cria um comando PowerShell para coletar informações sobre a placa-mãe.
// O comando retorna os dados em formato JSON, incluindo Manufacturer e Product.
//
// Retorna:
// - Um comando configurado para execução.
func (WindowsMotherboardRetriever) powerShellGetMotherboardInfo() *exec.Cmd {
	// Comando PowerShell para obter informações da placa-mãe em formato JSON
	cmd := exec.Command("powershell", "-Command", "Get-WmiObject -Class Win32_BaseBoard | Select-Object -Property Manufacturer, Product | ConvertTo-Json")
	return cmd
}

// deserializeMotherboardInfo desserializa os dados JSON em uma estrutura MotherboardInfo.
//
// Parâmetros:
// - data: Dados JSON a serem desserializados.
//
// Retorna:
// - Um objeto MotherboardInfo contendo as informações da placa-mãe.
// - Um erro, caso a desserialização falhe.
func (m WindowsMotherboardRetriever) deserializeMotherboardInfo(data []byte) (MotherboardInfo, error) {
	var motherboard MotherboardInfo
	err := json.Unmarshal(data, &motherboard)
	if err != nil {
		return MotherboardInfo{}, fmt.Errorf("erro ao deserializar JSON em MotherboardInfo: %v", err)
	}
	return motherboard, nil
}
