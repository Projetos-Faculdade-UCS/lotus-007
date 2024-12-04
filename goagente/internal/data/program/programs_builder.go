// Package programs fornece funcionalidades para coletar e estruturar informações sobre programas instalados no sistema.
// Este arquivo contém a definição de ProgramBuilder e ProgramInfo, usados para construir e armazenar essas informações.
package programs

// ProgramBuilder é o Builder para construir objetos ProgramInfo.
// Ele permite definir campos gradualmente e finalizar a construção com o método Build.
type ProgramBuilder struct {
	programInfo ProgramInfo
}

// ProgramInfo é a estrutura que armazena informações sobre programas instalados no sistema.
//
// Campos:
// - Patrimonio: Identificação patrimonial do sistema.
// - Programs: Lista de programas instalados no sistema.
// - HMAC: Hash criptográfico para validar a integridade dos dados (sempre no final).
type ProgramInfo struct {
	Patrimonio string    `json:"patrimonio"` // Identificação patrimonial
	Programs   []Program `json:"programs"`   // Lista de programas instalados
	HMAC       string    `json:"hmac"`       // Hash de integridade (sempre no final)
}

// SetPatrimonio define o campo Patrimonio no ProgramInfo.
//
// Parâmetros:
// - patrimonio: Identificação patrimonial do sistema.
//
// Retorna:
// - Uma referência ao próprio Builder.
func (b *ProgramBuilder) SetPatrimonio(patrimonio string) *ProgramBuilder {
	b.programInfo.Patrimonio = patrimonio
	return b
}

// SetPrograms define o campo Programs no ProgramInfo.
//
// Parâmetros:
// - programs: Lista de programas instalados no sistema.
//
// Retorna:
// - Uma referência ao próprio Builder.
func (b *ProgramBuilder) SetPrograms(programs []Program) *ProgramBuilder {
	b.programInfo.Programs = programs
	return b
}

// Build finaliza a construção do objeto ProgramInfo e o retorna.
//
// Retorna:
// - O objeto ProgramInfo totalmente construído.
func (b *ProgramBuilder) Build() ProgramInfo {
	return b.programInfo
}
