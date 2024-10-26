package orchestration

import (
	"goagente/internal/data/hardware"   // Importa as estruturas de dados como HardwareInfo
	"goagente/internal/data/patrimonio" // Para logar possíveis erros
	"goagente/internal/logging"
)

// HardwareOrchestrator é responsável por orquestrar a criação do objeto HardwareInfo
type HardwareOrchestrator struct{}

// NewHardwareOrchestrator cria uma nova instância do orquestrador de hardware
func NewHardwareOrchestrator() *HardwareOrchestrator {
	return &HardwareOrchestrator{}
}

// Orchestrate cria o objeto HardwareInfo agregando os dados de RAM, Disks, Processors e Motherboard
func (h *HardwareOrchestrator) Orchestrate() (hardware.HardwareInfo, error) {
	// Inicializa o builder para montar o objeto HardwareInfo
	builder := hardware.HardwareInfoBuilder{}

	// Chama os métodos automaticamente para preencher o hardware
	if err := h.autoPopulate(&builder); err != nil {
		return hardware.HardwareInfo{}, err
	}

	return builder.Build(), nil
}

// autoPopulate preenche automaticamente os dados de hardware no builder
func (h *HardwareOrchestrator) autoPopulate(builder *hardware.HardwareInfoBuilder) error {
	// Cria os retrievers automaticamente usando as fábricas
	patRetriever, _ := patrimonio.NewPatRetriever()
	pat, err := patRetriever.GetCurrentPat()
	if err != nil {
		logging.Error(err)
		return err
	}
	builder.SetPatrimonio(pat)

	ramRetriever, err := hardware.NewRAMRetriever()
	if err != nil {
		logging.Error(err)
		return err
	}

	diskRetriever, err := hardware.NewDiskRetriever()
	if err != nil {
		logging.Error(err)
		return err
	}

	processorRetriever, err := hardware.NewProcessorRetriever()
	if err != nil {
		logging.Error(err)
		return err
	}

	motherboardRetriever, err := hardware.NewMotherboardRetriever()
	if err != nil {
		logging.Error(err)
		return err
	}

	// Recupera as informações de RAM
	ramInfo, err := ramRetriever.GetRAMInfo()
	if err != nil {
		logging.Error(err)
		return err
	}
	builder.SetRAMModules(ramInfo)

	// Recupera as informações de discos
	diskInfo, err := diskRetriever.GetDiskInfo()
	if err != nil {
		logging.Error(err)
		return err
	}
	builder.SetDisks(diskInfo)

	// Recupera as informações de processadores
	processorInfo, err := processorRetriever.GetProcessorInfo()
	if err != nil {
		logging.Error(err)
		return err
	}
	builder.SetProcessors(processorInfo)

	// Recupera as informações da placa-mãe
	motherboardInfo, err := motherboardRetriever.GetMotherboardInfo()
	if err != nil {
		logging.Error(err)
		return err
	}
	builder.SetMotherboard(motherboardInfo)

	return nil
}
