package orchestration

import (
	"goagente/internal/communication"
	"goagente/internal/logging"
	"goagente/internal/security" // Importa a camada de segurança
	// Importa o pacote onde CoreInfoResult está definido
)

// CoreMediator é responsável por orquestrar e enviar informações de Core
type CoreMediator struct {
	poster    *communication.InfoPoster
	secretKey string // Adiciona a chave secreta para HMAC
}

// NewCoreMediator cria uma nova instância de CoreMediator com a chave secreta
func NewCoreMediator(poster *communication.InfoPoster, secretKey string) *CoreMediator {
	return &CoreMediator{
		poster:    poster,
		secretKey: secretKey, // Inicializa a chave secreta
	}
}

// OrchestrateAndPost coleta e envia as informações de Core para o servidor
func (m *CoreMediator) OrchestrateAndPost(patrimonio string) error {
	// Orquestra as informações de Core
	coreInfo, err := OrchestrateCoreInfo(patrimonio)
	if err != nil {
		logging.Error(err)
		return err
	}

	// Adiciona o HMAC diretamente na struct
	coreInfoWithHMAC, err := security.AddHMACToStruct(&coreInfo, m.secretKey)
	if err != nil {
		logging.Error(err)
		return err
	}

	// Envia o JSON com HMAC
	return m.poster.PostCoreInfo("core", coreInfoWithHMAC)
}
