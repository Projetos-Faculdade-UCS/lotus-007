package orchestration

import (
	"goagente/internal/data/patrimonio"
	"goagente/internal/data/system"
	"goagente/internal/logging"
	"log"
)

// OrchestrateCoreInfo preenche automaticamente o CoreInfoResult com hostname e usuário
func OrchestrateCoreInfo() (system.CoreInfoResult, error) {
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
	patRetriever, _ := patrimonio.NewPatRetriever()
	pat, err := patRetriever.GetCurrentPat()
	if err != nil {
		logging.Error(err)
		return system.CoreInfoResult{}, err
	}
	builder.SetPatrimonio(pat)

	//Coleta o sistema operacional
	osRetriever, err := system.NewOSRetriever()
	if err != nil {
		log.Println("Erro ao inicializar o OSRetriever:", err)
		return system.CoreInfoResult{}, err
	}
	os, err := osRetriever.GetCurrentOS()
	if err != nil {
		log.Println("Erro ao obter o sistema operacional:", err)
		return system.CoreInfoResult{}, err
	}
	builder.SetOs(os)

	return builder.Build(), nil
}
