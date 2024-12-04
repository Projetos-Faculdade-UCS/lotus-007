// Package orchestration gerencia a orquestração e coleta de informações sobre programas instalados e patrimônio.
// Este arquivo contém a implementação do ProgramOrchestrator, responsável por consolidar essas informações.
package orchestration

import (
	"goagente/internal/data/patrimonio"
	programs "goagente/internal/data/program"
	"goagente/internal/logging"
	"log"
)

// ProgramOrchestrator é responsável por orquestrar a criação do objeto ProgramInfo.
// Ele coleta informações sobre programas instalados e o patrimônio do sistema.
type ProgramOrchestrator struct{}

// NewProgramOrchestrator cria e retorna uma nova instância do ProgramOrchestrator.
//
// Retorna:
// - Uma instância de ProgramOrchestrator.
func NewProgramOrchestrator() *ProgramOrchestrator {
	return &ProgramOrchestrator{}
}

// Orchestrate cria o objeto ProgramInfo agregando os dados de programas instalados e patrimônio.
//
// Retorna:
// - Um objeto ProgramInfo preenchido com as informações coletadas.
// - Um erro, caso ocorra algum problema durante a coleta ou agregação de dados.
func (h *ProgramOrchestrator) Orchestrate() (programs.ProgramInfo, error) {
	// Inicializa o builder para montar o objeto ProgramInfo
	builder := programs.ProgramBuilder{}

	// Chama os métodos automaticamente para preencher o ProgramInfo
	if err := h.autoPopulate(&builder); err != nil {
		return programs.ProgramInfo{}, err
	}

	// Constrói e retorna o objeto final
	return builder.Build(), nil
}

// autoPopulate preenche automaticamente os dados de programas instalados e patrimônio no builder.
//
// Parâmetros:
// - builder: Um ponteiro para o ProgramBuilder que será preenchido com as informações.
//
// Retorna:
// - Um erro, caso ocorra algum problema durante a coleta de informações.
func (h *ProgramOrchestrator) autoPopulate(builder *programs.ProgramBuilder) error {
	// Coleta o patrimônio do sistema
	patRetriever, _ := patrimonio.NewPatRetriever()
	pat, err := patRetriever.GetCurrentPat()
	if err != nil {
		logging.Error(err)
		return err
	}
	builder.SetPatrimonio(pat)

	// Inicializa o ProgramRetriever para coletar programas instalados
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
