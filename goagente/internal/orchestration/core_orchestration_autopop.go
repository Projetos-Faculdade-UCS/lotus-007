// Package orchestration lida com a orquestração de informações e sua coleta de diversas fontes.
// Este arquivo contém a função OrchestrateCoreInfo, responsável por preencher o CoreInfoResult com dados do sistema.
package orchestration

import (
	"goagente/internal/data/patrimonio"
	"goagente/internal/data/system"
	"goagente/internal/logging"
	"log"
)

// OrchestrateCoreInfo coleta informações do sistema, como hostname, usuário atual, patrimônio e sistema operacional.
// Esses dados são utilizados para preencher e retornar um objeto do tipo CoreInfoResult.
//
// Retorna:
// - Um objeto CoreInfoResult contendo os dados coletados.
// - Um erro, caso ocorra algum problema durante a coleta de informações.
func OrchestrateCoreInfo() (system.CoreInfoResult, error) {
	// Inicializa o builder para construir o CoreInfoResult
	builder := system.CoreInfoResultBuilder{}

	// Coleta o hostname do sistema
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

	// Coleta o usuário atual do sistema
	userRetriever := system.WindowsUserRetriever{}
	username, err := userRetriever.GetCurrentUser()
	if err != nil {
		log.Println("Erro ao obter o usuário atual:", err)
		return system.CoreInfoResult{}, err
	}
	builder.SetUsername(username)

	// Coleta o patrimônio atual do sistema
	patRetriever, _ := patrimonio.NewPatRetriever()
	pat, err := patRetriever.GetCurrentPat()
	if err != nil {
		logging.Error(err)
		return system.CoreInfoResult{}, err
	}
	builder.SetPatrimonio(pat)

	// Coleta o sistema operacional do dispositivo
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

	// Constrói e retorna o CoreInfoResult
	return builder.Build(), nil
}
