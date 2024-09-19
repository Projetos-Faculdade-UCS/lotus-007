package orchestration

import (
	"goagente/internal/data/hardware" // Importa as estruturas de dados como HardwareInfo
	"goagente/internal/logging"       // Para logar possíveis erros
)

// HardwareOrchestrator é responsável por orquestrar a criação do objeto HardwareInfo
type HardwareOrchestrator struct {
	ramRetriever         hardware.RAMRetriever
	diskRetriever        hardware.DiskInfoRetriever
	processorRetriever   hardware.ProcessorInfoRetriever
	motherboardRetriever hardware.MotherboardInfoRetriever
}

// NewHardwareOrchestrator cria uma nova instância do orquestrador de hardware
func NewHardwareOrchestrator(
	ramRetriever hardware.RAMRetriever,
	diskRetriever hardware.DiskInfoRetriever,
	processorRetriever hardware.ProcessorInfoRetriever,
	motherboardRetriever hardware.MotherboardInfoRetriever,
) *HardwareOrchestrator {
	return &HardwareOrchestrator{
		ramRetriever:         ramRetriever,
		diskRetriever:        diskRetriever,
		processorRetriever:   processorRetriever,
		motherboardRetriever: motherboardRetriever,
	}
}

// Orchestrate cria o objeto HardwareInfo agregando os dados de RAM, Disks, Processors e Motherboard
func (h *HardwareOrchestrator) Orchestrate(patrimonio string) (hardware.HardwareInfo, error) {
	// Inicializa o builder para montar o objeto HardwareInfo
	builder := hardware.HardwareInfoBuilder{}

	// Recupera as informações de RAM
	ramInfo, err := h.ramRetriever.GetRAMInfo()
	if err != nil {
		logging.Error(err)
		return hardware.HardwareInfo{}, err
	}

	// Recupera as informações de discos
	diskInfo, err := h.diskRetriever.GetDiskInfo()
	if err != nil {
		logging.Error(err)
		return hardware.HardwareInfo{}, err
	}

	// Recupera as informações de processadores
	processorInfo, err := h.processorRetriever.GetProcessorInfo()
	if err != nil {
		logging.Error(err)
		return hardware.HardwareInfo{}, err
	}

	// Recupera as informações da placa-mãe
	motherboardInfo, err := h.motherboardRetriever.GetMotherboardInfo()
	if err != nil {
		logging.Error(err)
		return hardware.HardwareInfo{}, err
	}

	// Constrói o objeto HardwareInfo agregando todas as informações coletadas
	hardwareInfo := builder.SetPatrimonio(patrimonio).
		SetRAMModules(ramInfo).
		SetDisks(diskInfo).
		SetProcessors(processorInfo).
		SetMotherboard(motherboardInfo).
		Build()

	return hardwareInfo, nil
}
