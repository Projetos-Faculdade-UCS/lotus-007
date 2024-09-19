package hardware

import (
	"bytes"
	"encoding/json"
	"fmt"
	"goagente/pkg/utils"
	"os/exec"
)

type LinuxDiskRetriever struct{}

func (g LinuxDiskRetriever) GetDiskInfo() ([]DiskInfo, error) {
	cmd := exec.Command("lsblk", "-J", "-o", "NAME,SIZE,MODEL")

	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return nil, fmt.Errorf("erro ao executar o comando lsblk: %v", err)
	}

	// Process JSON output here, similar to WindowsDiskRetriever
	var disks []DiskInfo
	err = json.Unmarshal(out.Bytes(), &disks)
	if err != nil {
		return nil, fmt.Errorf("erro ao deserializar o JSON do lsblk: %v", err)
	}

	// Convert size to gigabytes (assuming the output is in bytes)
	for i := range disks {
		disks[i].Size = uint64(utils.BytesToGigabytes(disks[i].Size))
	}

	return disks, nil
}
