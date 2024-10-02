package hardware

import (
	"bytes"
	"fmt"
	"os/exec"
)

type LinuxRAMRetriever struct{}

func (l LinuxRAMRetriever) GetRAMInfo() ([]RAM, error) {
	cmd := exec.Command("dmidecode", "--type", "17") // Gets RAM info using dmidecode on Linux

	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return nil, fmt.Errorf("error executing dmidecode: %v", err)
	}

	// Process dmidecode output
	ramList, err := l.parseDMIDecodeOutput(out.Bytes())
	if err != nil {
		return nil, err
	}

	return ramList, nil
}

// parseDMIDecodeOutput parses the output of the dmidecode command
func (l LinuxRAMRetriever) parseDMIDecodeOutput(data []byte) ([]RAM, error) {
	ramList := []RAM{}
	lines := bytes.Split(data, []byte("\n"))

	var currentRAM RAM
	for _, line := range lines {
		trimmedLine := bytes.TrimSpace(line)

		// Verifica o campo "Manufacturer"
		if bytes.HasPrefix(trimmedLine, []byte("Manufacturer:")) {
			manufacturer := string(bytes.TrimPrefix(trimmedLine, []byte("Manufacturer:")))
			currentRAM.Manufacturer = manufacturer
		}

		// Verifica o campo "Size" e converte para GB
		if bytes.HasPrefix(trimmedLine, []byte("Size:")) {
			sizeStr := string(bytes.TrimPrefix(trimmedLine, []byte("Size:")))
			if sizeStr != "No Module Installed" {
				sizeBytes, err := parseSize(sizeStr)
				if err != nil {
					return nil, fmt.Errorf("erro ao processar tamanho da RAM: %v", err)
				}
				// Converte de bytes para GB
				currentRAM.Capacity = float64(sizeBytes) / (1024 * 1024 * 1024)
			}
		}

		// Verifica o campo "Form Factor"
		if bytes.HasPrefix(trimmedLine, []byte("Form Factor:")) {
			formFactorStr := string(bytes.TrimPrefix(trimmedLine, []byte("Form Factor:")))
			currentRAM.FormFactor = parseFormFactor(formFactorStr)
		}

		// Quando encontrar o fim de uma entrada de RAM, adicione à lista e resete currentRAM
		if bytes.Equal(trimmedLine, []byte("")) && currentRAM.Capacity != 0 {
			ramList = append(ramList, currentRAM)
			currentRAM = RAM{} // Reseta para a próxima entrada
		}
	}

	return ramList, nil
}

// parseSize converte o tamanho da RAM de string para bytes
func parseSize(sizeStr string) (uint64, error) {
	// Remove " MB" ou " GB" e converte para número
	var size uint64
	var unit string
	_, err := fmt.Sscanf(sizeStr, "%d %s", &size, &unit)
	if err != nil {
		return 0, err
	}

	switch unit {
	case "MB":
		return size * 1024 * 1024, nil
	case "GB":
		return size * 1024 * 1024 * 1024, nil
	default:
		return 0, fmt.Errorf("unidade desconhecida: %s", unit)
	}
}

// parseFormFactor converte a string do Form Factor em um número
func parseFormFactor(formFactorStr string) int {
	switch formFactorStr {
	case "DIMM":
		return 8
	case "SO-DIMM":
		return 9
	default:
		return 0 // Valor desconhecido ou não especificado
	}
}
