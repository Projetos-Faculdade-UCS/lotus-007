// Package orchestration lida com a orquestração de informações e sua comunicação com outros componentes.
// Este arquivo contém a implementação do HardwareMediator, responsável por orquestrar e enviar informações de hardware.
package orchestration

import (
	"goagente/internal/communication"
	"goagente/internal/logging"
	"goagente/internal/security"
)

// HardwareMediator é responsável por orquestrar e enviar informações de hardware.
// Ele utiliza um InfoPoster para enviar dados e uma chave secreta para operações de HMAC.
type HardwareMediator struct {
	orchestrator *HardwareOrchestrator     // Orquestrador que coleta informações de hardware
	poster       *communication.InfoPoster // Responsável por enviar as informações para o servidor
	secretKey    string                    // Chave secreta usada para adicionar HMAC às informações
}

// NewHardwareMediator cria e retorna uma nova instância de HardwareMediator.
//
// Parâmetros:
// - orchestrator: Instância de HardwareOrchestrator utilizada para coletar informações de hardware.
// - poster: Instância de InfoPoster utilizada para enviar as informações ao servidor.
// - secretKey: Chave secreta usada para adicionar HMAC às informações.
//
// Retorna:
// - Uma nova instância de HardwareMediator.
func NewHardwareMediator(orchestrator *HardwareOrchestrator, poster *communication.InfoPoster, secretKey string) *HardwareMediator {
	return &HardwareMediator{
		orchestrator: orchestrator,
		poster:       poster,
		secretKey:    secretKey, // Inicializa a chave secreta
	}
}

// OrchestrateAndPost coleta as informações de hardware, adiciona um HMAC e as envia ao servidor.
//
// Parâmetros:
// - param: Parâmetro utilizado durante a orquestração das informações (reservado para uso futuro ou opcional).
//
// Retorna:
// - Um erro, caso ocorra algum problema na orquestração, adição do HMAC ou envio das informações.
func (m *HardwareMediator) OrchestrateAndPost(param string) error {

	// Orquestra as informações de hardware
	hardwareInfo, err := m.orchestrator.Orchestrate()
	if err != nil {
		logging.Error(err) // Registra o erro no sistema de logging
		return err
	}

	// Adiciona o HMAC diretamente na estrutura de hardwareInfo
	hardwareInfoWithHMAC, err := security.AddHMACToStruct(&hardwareInfo, m.secretKey)
	if err != nil {
		logging.Error(err) // Registra o erro no sistema de logging
		return err
	}

	// Envia as informações de hardware com o HMAC para o servidor
	return m.poster.PostHardwareInfo("agente/hardware/", hardwareInfoWithHMAC)
}
