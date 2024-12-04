// Package hardware fornece funcionalidades para coletar informações sobre discos físicos do sistema.
// Este arquivo contém a implementação de WindowsDiskRetriever para coletar dados no sistema operacional Windows.
package hardware

import (
	"bytes"
	"encoding/json"
	"fmt"
	"goagente/internal/logging"
	"goagente/pkg/utils"
	"os/exec"
)

// WindowsDiskRetriever é a implementação de DiskInfoRetriever para o sistema operacional Windows.
// Ele utiliza comandos PowerShell para coletar informações sobre discos físicos.
type WindowsDiskRetriever struct{}

// GetDiskInfo coleta informações sobre discos físicos no Windows.
// Ele utiliza o PowerShell para obter os dados e converte o tamanho dos discos de bytes para gigabytes.
//
// Retorna:
// - Um slice de DiskInfo contendo as informações dos discos.
// - Um erro, caso ocorra algum problema durante a execução do comando ou a desserialização dos dados.
func (d WindowsDiskRetriever) GetDiskInfo() ([]DiskInfo, error) {
	// Executa o comando PowerShell para obter informações sobre discos
	cmd := d.powerShellGetDiskInfo()

	// Captura a saída do comando
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		newErr := fmt.Errorf("erro ao executar o comando PowerShell em getdiskinfo: %v", err)
		logging.Error(newErr)
		return nil, err
	}

	// Desserializa a saída JSON em objetos DiskInfo
	disks, err := d.deserializeDiskInfo(out.Bytes())
	if err != nil {
		logging.Error(err)
		return nil, err
	}

	// Converte o tamanho dos discos para gigabytes
	for i := range disks {
		disks[i].Size = uint64(utils.BytesToGigabytes(disks[i].Size))
	}

	return disks, nil
}

// powerShellGetDiskInfo cria um comando PowerShell para coletar informações sobre discos físicos.
// O comando retorna os dados em formato JSON, incluindo DeviceID, Model e Size.
//
// Retorna:
// - Um comando configurado para execução.
func (WindowsDiskRetriever) powerShellGetDiskInfo() *exec.Cmd {
	// Comando PowerShell para obter informações do disco em formato JSON
	cmd := exec.Command("powershell", "-Command", "Get-PhysicalDisk | Select-Object -Property DeviceID, Model, Size | ConvertTo-Json")
	return cmd
}

// deserializeDiskInfo tenta desserializar os dados JSON em um ou mais objetos DiskInfo.
//
// Funcionalidade:
// - Primeiro, tenta desserializar como um único objeto DiskInfo.
// - Se falhar, tenta desserializar como um array de objetos DiskInfo.
//
// Parâmetros:
// - data: Dados JSON a serem desserializados.
//
// Retorna:
// - Um slice de DiskInfo contendo as informações dos discos.
// - Um erro, caso a desserialização falhe.
func (d WindowsDiskRetriever) deserializeDiskInfo(data []byte) ([]DiskInfo, error) {
	// Tenta desserializar o JSON como um único objeto
	var singleDisk DiskInfo
	err := json.Unmarshal(data, &singleDisk)
	if err == nil {
		// Se for um único objeto, retorna como um slice com um único item
		return []DiskInfo{singleDisk}, nil
	}

	// Se falhar, tenta deserializar como um array de objetos
	var disks []DiskInfo
	err = json.Unmarshal(data, &disks)
	if err != nil {
		return nil, fmt.Errorf("erro ao deserializar JSON em diskInfo: %v", err)
	}

	return disks, nil
}
