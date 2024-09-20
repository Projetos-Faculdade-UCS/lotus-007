package orchestration

import (
	"goagente/internal/data/system"
	"log"
)

// OrchestrateCoreInfo preenche automaticamente o CoreInfoResult com hostname e usuário
func OrchestrateCoreInfo(patrimonio string) (system.CoreInfoResult, error) {
	// Inicializa o builder
	builder := system.CoreInfoResultBuilder{}

	// Coleta o hostname
	hostnameRetriever, err := system.NewHostnameRetriever()
	if err != nil {
		log.Println("Erro ao inicializar o HostnameRetriever:", err)
		return system.CoreInfoResult{}, err
	}
	hostname, err := hostnameRetriever.GetHostname()
	if err != nil {
		log.Println("Erro ao obter o hostname:", err)
		return system.CoreInfoResult{}, err
	}
	builder.SetHostname(hostname)

	// Coleta o usuário atual
	userRetriever := system.WindowsUserRetriever{}
	username, err := userRetriever.GetCurrentUser()
	if err != nil {
		log.Println("Erro ao obter o usuário atual:", err)
		return system.CoreInfoResult{}, err
	}
	builder.SetUsername(username)

	// Define o patrimônio e constrói o objeto CoreInfoResult
	builder.SetPatrimonio(patrimonio)
	return builder.Build(), nil
}
