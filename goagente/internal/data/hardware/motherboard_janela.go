package hardware

import (
	"bytes"
	"encoding/json"
	"fmt"
	"goagente/internal/logging"
	"os/exec"
)

type WindowsMotherboardRetriever struct{}

func (r WindowsMotherboardRetriever) GetMotherboardInfo() (MotherboardInfo, error) {
	cmd := r.powerShellGetMotherboardInfo()

	// Executa o comando PowerShell
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		newErr := fmt.Errorf("erro ao executar o comando PowerShell em GetMotherboardInfo: %v", err)
		logging.Error(newErr)
		return MotherboardInfo{}, err
	}

	// Usa o novo método para desserializar o JSON
	motherboard, err := r.deserializeMotherboardInfo(out.Bytes())
	if err != nil {
		logging.Error(err)
		return MotherboardInfo{}, err
	}

	return motherboard, nil
}

// PowerShellGetMotherboardInfo executa o comando PowerShell para obter informações da placa-mãe
func (WindowsMotherboardRetriever) powerShellGetMotherboardInfo() *exec.Cmd {
	// Comando PowerShell para obter informações da placa-mãe em formato JSON
	cmd := exec.Command("powershell", "-Command", "Get-WmiObject -Class Win32_BaseBoard | Select-Object -Property Manufacturer, Product | ConvertTo-Json")
	return cmd
}

// deserializeMotherboardInfo desserializa o JSON para a estrutura MotherboardInfo
func (m WindowsMotherboardRetriever) deserializeMotherboardInfo(data []byte) (MotherboardInfo, error) {
	var motherboard MotherboardInfo
	err := json.Unmarshal(data, &motherboard)
	if err != nil {
		return MotherboardInfo{}, fmt.Errorf("erro ao deserializar JSON em MotherboardInfo: %v", err)
	}
	return motherboard, nil
}
