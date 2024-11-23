// Package orchestration lida com a orquestração de informações e sua comunicação com outros componentes.
// Este arquivo contém a implementação do CoreMediator, responsável por orquestrar e enviar informações de Core.
package orchestration

import (
	"goagente/internal/communication" // Camada de comunicação com o servidor
	"goagente/internal/logging"       // Camada de logging
	"goagente/internal/security"      // Camada de segurança para adicionar HMAC
)

// CoreMediator é responsável por orquestrar e enviar informações de Core.
// Ele utiliza um InfoPoster para enviar dados e uma chave secreta para operações de HMAC.
type CoreMediator struct {
	poster    *communication.InfoPoster // Responsável por enviar as informações para o servidor
	secretKey string                    // Chave secreta usada para adicionar HMAC às informações
}

// NewCoreMediator cria e retorna uma nova instância de CoreMediator.
//
// Parâmetros:
// - poster: Instância de InfoPoster utilizada para enviar as informações ao servidor.
// - secretKey: Chave secreta usada para adicionar HMAC às informações.
//
// Retorna:
// - Uma nova instância de CoreMediator.
func NewCoreMediator(poster *communication.InfoPoster, secretKey string) *CoreMediator {
	return &CoreMediator{
		poster:    poster,
		secretKey: secretKey, // Inicializa a chave secreta
	}
}

// OrchestrateAndPost coleta as informações de Core, adiciona um HMAC, e as envia ao servidor.
//
// Parâmetros:
// - param: Parâmetro utilizado durante a orquestração das informações (futuro ou opcional).
//
// Retorna:
// - Um erro, caso ocorra algum problema na orquestração, adição do HMAC ou envio das informações.
func (m *CoreMediator) OrchestrateAndPost(param string) error {
	// Orquestra as informações de Core
	coreInfo, err := OrchestrateCoreInfo()
	if err != nil {
		logging.Error(err) // Registra o erro no sistema de logging
		return err
	}

	// Adiciona o HMAC diretamente na estrutura de CoreInfo
	coreInfoWithHMAC, err := security.AddHMACToStruct(&coreInfo, m.secretKey)
	if err != nil {
		logging.Error(err) // Registra o erro no sistema de logging
		return err
	}

	// Envia as informações de Core com o HMAC para o servidor
	return m.poster.PostCoreInfo("agente/core/", coreInfoWithHMAC)
}
