package system

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
