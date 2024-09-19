package system

import (
	"log"
)

// CoreInfoResult é a estrutura que queremos construir
type CoreInfoResult struct {
	Patrimonio string `json:"patrimonio"`
	Hostname   string `json:"hostname"`
	Username   string `json:"username"`
}

// CoreInfoResultBuilder é o Builder para construir CoreInfoResult
type CoreInfoResultBuilder struct {
	coreInfo CoreInfoResult
}

// SetHostname define o Hostname no CoreInfoResult
func (b *CoreInfoResultBuilder) SetHostname(hostname string) *CoreInfoResultBuilder {
	b.coreInfo.Hostname = hostname
	return b
}

// SetUsername define o Username no CoreInfoResult
func (b *CoreInfoResultBuilder) SetUsername(username string) *CoreInfoResultBuilder {
	b.coreInfo.Username = username
	return b
}

// SetPatrimonio define o Patrimonio no CoreInfoResult
func (b *CoreInfoResultBuilder) SetPatrimonio(patrimonio string) *CoreInfoResultBuilder {
	b.coreInfo.Patrimonio = patrimonio
	return b
}

// Build finaliza a construção e retorna o CoreInfoResult
func (b *CoreInfoResultBuilder) Build() CoreInfoResult {
	return b.coreInfo
}

// AutomaticPopulate preenche automaticamente o CoreInfoResult com hostname e usuário
func (b *CoreInfoResultBuilder) AutomaticPopulate() (*CoreInfoResultBuilder, error) {
	// Coleta o hostname
	hostnameRetriever, err := NewHostnameRetriever()
	if err != nil {
		log.Println("Erro ao inicializar o HostnameRetriever:", err)
		return b, err
	}
	hostname, err := hostnameRetriever.GetHostname()
	if err != nil {
		log.Println("Erro ao obter o hostname:", err)
		return b, err
	}
	b.SetHostname(hostname)

	// Coleta o usuário atual no Windows
	userRetriever := WindowsUserRetriever{}
	username, err := userRetriever.GetCurrentUser()
	if err != nil {
		log.Println("Erro ao obter o usuário atual:", err)
		return b, err
	}
	b.SetUsername(username)

	return b, nil
}
