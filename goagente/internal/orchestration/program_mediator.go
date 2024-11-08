package orchestration

import (
	"goagente/internal/communication"
	"goagente/internal/logging"
	"goagente/internal/security" // Importa a camada de segurança
)

// ProgramMediator é responsável por orquestrar e enviar informações de programas instalados
type ProgramMediator struct {
	orchestrator *ProgramOrchestrator
	poster       *communication.InfoPoster
	secretKey    string // Adiciona a chave secreta para HMAC
}

// NewProgramMediator cria uma nova instância de ProgramMediator com a chave secreta
func NewProgramMediator(orchestrator *ProgramOrchestrator, poster *communication.InfoPoster, secretKey string) *ProgramMediator {
	return &ProgramMediator{
		orchestrator: orchestrator,
		poster:       poster,
		secretKey:    secretKey, // Inicializa a chave secreta
	}
}

// OrchestrateAndPost coleta e envia as informações de programas instalados para o servidor
func (m *ProgramMediator) OrchestrateAndPost(param string) error {
	programInfo, err := m.orchestrator.Orchestrate()
	if err != nil {
		logging.Error(err)
		return err
	}

	// Adiciona o HMAC à struct antes de serializar
	programInfoWithHMAC, err := security.AddHMACToStruct(&programInfo, m.secretKey)
	if err != nil {
		logging.Error(err)
		return err
	}

	// Envia o JSON com HMAC
	return m.poster.PostProgramInfo("v3/f7152de8-232a-440f-8a51-8c453c9356ac", programInfoWithHMAC)
}
