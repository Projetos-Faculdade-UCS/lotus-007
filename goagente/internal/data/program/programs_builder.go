package programs

// ProgramBuilder é o Builder para construir a lista de programas com patrimônio
type ProgramBuilder struct {
	programInfo ProgramInfo
}

// ProgramInfo é a estrutura que será criada pelo Builder
type ProgramInfo struct {
	Patrimonio string    `json:"patrimonio"`
	Programs   []Program `json:"programs"`
}

// SetPatrimonio define o patrimônio no ProgramInfo
func (b *ProgramBuilder) SetPatrimonio(patrimonio string) *ProgramBuilder {
	b.programInfo.Patrimonio = patrimonio
	return b
}

// SetPrograms define a lista de programas instalados no ProgramInfo
func (b *ProgramBuilder) SetPrograms(programs []Program) *ProgramBuilder {
	b.programInfo.Programs = programs
	return b
}

// Build finaliza a construção e retorna o ProgramInfo
func (b *ProgramBuilder) Build() ProgramInfo {
	return b.programInfo
}
