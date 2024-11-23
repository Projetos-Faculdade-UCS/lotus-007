// Package orchestration gerencia a coleta, processamento e envio de informações de programas instalados.
// Este arquivo contém a implementação do ProgramMediator, responsável por orquestrar e enviar essas informações.
package orchestration

import (
	"goagente/internal/communication"
	"goagente/internal/logging"
	"goagente/internal/security"
)

// ProgramMediator é responsável por orquestrar e enviar informações de programas instalados.
// Ele utiliza um ProgramOrchestrator para coletar os dados, adiciona um HMAC para segurança e envia as informações ao servidor.
type ProgramMediator struct {
	orchestrator *ProgramOrchestrator      // Orquestrador que coleta informações de programas instalados
	poster       *communication.InfoPoster // Ferramenta para envio de informações ao servidor
	secretKey    string                    // Chave secreta usada para adicionar HMAC às informações
}

// NewProgramMediator cria e retorna uma nova instância de ProgramMediator.
//
// Parâmetros:
// - orchestrator: Instância de ProgramOrchestrator utilizada para coletar informações de programas instalados.
// - poster: Instância de InfoPoster utilizada para enviar as informações ao servidor.
// - secretKey: Chave secreta usada para adicionar HMAC às informações.
//
// Retorna:
// - Uma nova instância de ProgramMediator.
func NewProgramMediator(orchestrator *ProgramOrchestrator, poster *communication.InfoPoster, secretKey string) *ProgramMediator {
	return &ProgramMediator{
		orchestrator: orchestrator,
		poster:       poster,
		secretKey:    secretKey, // Inicializa a chave secreta
	}
}

// OrchestrateAndPost coleta informações de programas instalados, adiciona um HMAC para segurança e as envia ao servidor.
//
// Parâmetros:
// - param: Parâmetro opcional para uso futuro, atualmente não utilizado.
//
// Retorna:
// - Um erro, caso ocorra algum problema na coleta, adição do HMAC ou envio das informações.
func (m *ProgramMediator) OrchestrateAndPost(param string) error {
	// Coleta informações de programas instalados
	programInfo, err := m.orchestrator.Orchestrate()
	if err != nil {
		logging.Error(err) // Registra o erro no sistema de logging
		return err
	}

	// Adiciona o HMAC à struct antes de serializar
	programInfoWithHMAC, err := security.AddHMACToStruct(&programInfo, m.secretKey)
	if err != nil {
		logging.Error(err) // Registra o erro no sistema de logging
		return err
	}

	// Envia as informações de programas instalados com o HMAC ao servidor
	return m.poster.PostProgramInfo("agente/programs/", programInfoWithHMAC)
}
