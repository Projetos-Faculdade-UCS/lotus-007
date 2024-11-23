// Package hardware fornece funcionalidades para coletar e estruturar informações de hardware do sistema.
// Este arquivo contém a definição de HardwareInfo e o padrão Builder para facilitar a construção do objeto.
package hardware

// HardwareInfo representa informações completas sobre o hardware do sistema.
//
// Campos:
// - Patrimonio: Identificação patrimonial do sistema.
// - Disks: Lista de informações sobre os discos físicos.
// - Processors: Lista de informações sobre os processadores.
// - RAMModules: Lista de informações sobre os módulos de memória RAM.
// - Motherboard: Informações sobre a placa-mãe.
// - HMAC: Hash criptográfico para validar a integridade dos dados (sempre no final).
type HardwareInfo struct {
	Patrimonio  string          `json:"patrimonio"`  // Identificação patrimonial
	Disks       []DiskInfo      `json:"disks"`       // Informações dos discos
	Processors  []ProcessorInfo `json:"processors"`  // Informações dos processadores
	RAMModules  []RAM           `json:"ram"`         // Informações dos módulos de RAM
	Motherboard MotherboardInfo `json:"motherboard"` // Informações da placa-mãe
	HMAC        string          `json:"hmac"`        // Hash de integridade (sempre no final)
}

// HardwareInfoBuilder fornece um padrão Builder para construir objetos HardwareInfo.
// Ele permite definir os campos gradualmente e finalizar a construção com o método Build.
type HardwareInfoBuilder struct {
	hardware HardwareInfo
}

// SetPatrimonio define o campo Patrimonio no objeto HardwareInfo.
//
// Parâmetros:
// - patrimonio: Identificação patrimonial do sistema.
//
// Retorna:
// - Uma referência ao próprio Builder.
func (b *HardwareInfoBuilder) SetPatrimonio(patrimonio string) *HardwareInfoBuilder {
	b.hardware.Patrimonio = patrimonio
	return b
}

// SetDisks define o campo Disks no objeto HardwareInfo.
//
// Parâmetros:
// - disks: Lista de informações sobre os discos físicos.
//
// Retorna:
// - Uma referência ao próprio Builder.
func (b *HardwareInfoBuilder) SetDisks(disks []DiskInfo) *HardwareInfoBuilder {
	b.hardware.Disks = disks
	return b
}

// SetProcessors define o campo Processors no objeto HardwareInfo.
//
// Parâmetros:
// - processors: Lista de informações sobre os processadores.
//
// Retorna:
// - Uma referência ao próprio Builder.
func (b *HardwareInfoBuilder) SetProcessors(processors []ProcessorInfo) *HardwareInfoBuilder {
	b.hardware.Processors = processors
	return b
}

// SetRAMModules define o campo RAMModules no objeto HardwareInfo.
//
// Parâmetros:
// - ram: Lista de informações sobre os módulos de memória RAM.
//
// Retorna:
// - Uma referência ao próprio Builder.
func (b *HardwareInfoBuilder) SetRAMModules(ram []RAM) *HardwareInfoBuilder {
	b.hardware.RAMModules = ram
	return b
}

// SetMotherboard define o campo Motherboard no objeto HardwareInfo.
//
// Parâmetros:
// - motherboard: Informações sobre a placa-mãe.
//
// Retorna:
// - Uma referência ao próprio Builder.
func (b *HardwareInfoBuilder) SetMotherboard(motherboard MotherboardInfo) *HardwareInfoBuilder {
	b.hardware.Motherboard = motherboard
	return b
}

// Build finaliza a construção do objeto HardwareInfo e o retorna.
//
// Retorna:
// - O objeto HardwareInfo totalmente construído.
func (b *HardwareInfoBuilder) Build() HardwareInfo {
	return b.hardware
}
