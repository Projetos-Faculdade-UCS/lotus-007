package orchestration

import (
	"goagente/internal/data/patrimonio"
	programs "goagente/internal/data/program"
	"goagente/internal/logging"
	"log"
)

// ProgramOrchestrator é responsável por orquestrar a criação do objeto ProgramInfo
type ProgramOrchestrator struct{}

// NewProgramOrchestrator cria uma nova instância do orquestrador de programas
func NewProgramOrchestrator() *ProgramOrchestrator {
	return &ProgramOrchestrator{}
}

// Orchestrate cria o objeto ProgramInfo agregando os dados de programas instalados e patrimônio
func (h *ProgramOrchestrator) Orchestrate() (programs.ProgramInfo, error) {
	// Inicializa o builder para montar o objeto ProgramInfo
	builder := programs.ProgramBuilder{}

	// Chama os métodos automaticamente para preencher o ProgramInfo
	if err := h.autoPopulate(&builder); err != nil {
		return programs.ProgramInfo{}, err
	}

	// Define o patrimônio e constrói o objeto final
	return builder.Build(), nil
}

// autoPopulate preenche automaticamente os dados de programas instalados no builder
func (h *ProgramOrchestrator) autoPopulate(builder *programs.ProgramBuilder) error {

	patRetriever, _ := patrimonio.NewPatRetriever()
	pat, err := patRetriever.GetCurrentPat()
	if err != nil {
		logging.Error(err)
		return err
	}
	builder.SetPatrimonio(pat)

	programRetriever, err := programs.NewProgramRetriever()
	if err != nil {
		log.Println("Erro ao inicializar o ProgramRetriever:", err)
		return err
	}

	// Recupera a lista de programas instalados
	programsInstalled, err := programRetriever.GetInstalledPrograms()
	if err != nil {
		log.Println("Erro ao recuperar a lista de programas instalados:", err)
		return err
	}

	// Define os programas instalados no builder
	builder.SetPrograms(programsInstalled)

	return nil
}
