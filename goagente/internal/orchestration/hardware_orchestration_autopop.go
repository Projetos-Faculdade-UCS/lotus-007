// Package orchestration lida com a orquestração de informações de hardware e sua organização em objetos.
// Este arquivo contém a implementação do HardwareOrchestrator, responsável por coletar e montar os dados de hardware.
package orchestration

import (
	"goagente/internal/data/hardware"   // Importa as estruturas de dados relacionadas ao hardware
	"goagente/internal/data/patrimonio" // Importa as funcionalidades para tratar informações de patrimônio
	"goagente/internal/logging"         // Importa a funcionalidade de logging
)

// HardwareOrchestrator é responsável por orquestrar a criação do objeto HardwareInfo.
// Ele coleta informações de RAM, discos, processadores e placa-mãe, consolidando-as em um único objeto.
type HardwareOrchestrator struct{}

// NewHardwareOrchestrator cria e retorna uma nova instância de HardwareOrchestrator.
//
// Retorna:
// - Uma instância de HardwareOrchestrator.
func NewHardwareOrchestrator() *HardwareOrchestrator {
	return &HardwareOrchestrator{}
}

// Orchestrate cria o objeto HardwareInfo, agregando os dados de RAM, discos, processadores e placa-mãe.
//
// Retorna:
// - Um objeto HardwareInfo preenchido com as informações coletadas.
// - Um erro, caso ocorra algum problema na coleta dos dados.
func (h *HardwareOrchestrator) Orchestrate() (hardware.HardwareInfo, error) {
	// Inicializa o builder para montar o objeto HardwareInfo
	builder := hardware.HardwareInfoBuilder{}

	// Preenche os dados de hardware automaticamente no builder
	if err := h.autoPopulate(&builder); err != nil {
		return hardware.HardwareInfo{}, err
	}

	// Constrói e retorna o objeto HardwareInfo
	return builder.Build(), nil
}

// autoPopulate preenche automaticamente os dados de hardware no builder.
// Ele utiliza diferentes retrievers para coletar informações de RAM, discos, processadores e placa-mãe.
//
// Parâmetros:
// - builder: Um ponteiro para o builder que será preenchido com as informações.
//
// Retorna:
// - Um erro, caso ocorra algum problema na coleta de informações.
func (h *HardwareOrchestrator) autoPopulate(builder *hardware.HardwareInfoBuilder) error {
	// Coleta as informações de patrimônio
	patRetriever, _ := patrimonio.NewPatRetriever()
	pat, err := patRetriever.GetCurrentPat()
	if err != nil {
		logging.Error(err)
		return err
	}
	builder.SetPatrimonio(pat)

	// Coleta as informações de RAM
	ramRetriever, err := hardware.NewRAMRetriever()
	if err != nil {
		logging.Error(err)
		return err
	}
	ramInfo, err := ramRetriever.GetRAMInfo()
	if err != nil {
		logging.Error(err)
		return err
	}
	builder.SetRAMModules(ramInfo)

	// Coleta as informações de discos
	diskRetriever, err := hardware.NewDiskRetriever()
	if err != nil {
		logging.Error(err)
		return err
	}
	diskInfo, err := diskRetriever.GetDiskInfo()
	if err != nil {
		logging.Error(err)
		return err
	}
	builder.SetDisks(diskInfo)

	// Coleta as informações de processadores
	processorRetriever, err := hardware.NewProcessorRetriever()
	if err != nil {
		logging.Error(err)
		return err
	}
	processorInfo, err := processorRetriever.GetProcessorInfo()
	if err != nil {
		logging.Error(err)
		return err
	}
	builder.SetProcessors(processorInfo)

	// Coleta as informações da placa-mãe
	motherboardRetriever, err := hardware.NewMotherboardRetriever()
	if err != nil {
		logging.Error(err)
		return err
	}
	motherboardInfo, err := motherboardRetriever.GetMotherboardInfo()
	if err != nil {
		logging.Error(err)
		return err
	}
	builder.SetMotherboard(motherboardInfo)

	return nil
}
