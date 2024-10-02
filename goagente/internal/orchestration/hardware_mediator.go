package orchestration

import (
	"goagente/internal/communication"
	"goagente/internal/logging"
	"goagente/internal/security" // Importa o pacote de segurança
)

// HardwareMediator é responsável por orquestrar e enviar informações de hardware
type HardwareMediator struct {
	orchestrator *HardwareOrchestrator
	poster       *communication.InfoPoster
	secretKey    string // Adiciona a chave secreta
}

// NewHardwareMediator cria uma nova instância de HardwareMediator
func NewHardwareMediator(orchestrator *HardwareOrchestrator, poster *communication.InfoPoster, secretKey string) *HardwareMediator {
	return &HardwareMediator{
		orchestrator: orchestrator,
		poster:       poster,
		secretKey:    secretKey, // Inicializa a chave secreta
	}
}

// OrchestrateAndPost coleta e envia as informações de hardware para o servidor
func (m *HardwareMediator) OrchestrateAndPost(patrimonio string) error {
	// Orquestra as informações de hardware
	hardwareInfo, err := m.orchestrator.Orchestrate(patrimonio)
	if err != nil {
		logging.Error(err)
		return err
	}

	// Adiciona o HMAC diretamente na struct
	hardwareInfoWithHMAC, err := security.AddHMACToStruct(&hardwareInfo, m.secretKey)
	if err != nil {
		logging.Error(err)
		return err
	}

	// Envia o JSON com HMAC
	return m.poster.PostHardwareInfo("hardware", hardwareInfoWithHMAC)
}
