package hardware

import (
	"bytes"
	"fmt"
	"os/exec"
)

type LinuxMotherboardRetriever struct{}

func (r LinuxMotherboardRetriever) GetMotherboardInfo() (MotherboardInfo, error) {
	cmd := exec.Command("sh", "-c", "cat /sys/class/dmi/id/board_vendor && cat /sys/class/dmi/id/board_name")

	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return MotherboardInfo{}, fmt.Errorf("erro ao executar o comando para obter informações da placa-mãe no Linux: %v", err)
	}

	// Parsing output for Linux (assumindo que o comando retorna duas linhas: fabricante e modelo)
	lines := bytes.Split(out.Bytes(), []byte("\n"))
	if len(lines) < 2 {
		return MotherboardInfo{}, fmt.Errorf("erro ao obter informações da placa-mãe no Linux: saída inesperada")
	}

	motherboard := MotherboardInfo{
		Manufacturer: string(lines[0]),
		Product:      string(lines[1]),
	}
	return motherboard, nil
}
