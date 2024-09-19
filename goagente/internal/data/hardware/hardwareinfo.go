package hardware

type HardwareInfo struct {
	Patrimonio  string          `json:"patrimonio"`
	Disks       []DiskInfo      `json:"disks"`
	Processors  []ProcessorInfo `json:"processors"`
	RAMModules  []RAM           `json:"ram"`
	Motherboard MotherboardInfo `json:"motherboard"`
}

type HardwareInfoBuilder struct {
	hardware HardwareInfo
}

func (b *HardwareInfoBuilder) SetPatrimonio(patrimonio string) *HardwareInfoBuilder {
	b.hardware.Patrimonio = patrimonio
	return b
}

func (b *HardwareInfoBuilder) SetDisks(disks []DiskInfo) *HardwareInfoBuilder {
	b.hardware.Disks = disks
	return b
}

func (b *HardwareInfoBuilder) SetProcessors(processors []ProcessorInfo) *HardwareInfoBuilder {
	b.hardware.Processors = processors
	return b
}

func (b *HardwareInfoBuilder) SetRAMModules(ram []RAM) *HardwareInfoBuilder {
	b.hardware.RAMModules = ram
	return b
}

func (b *HardwareInfoBuilder) SetMotherboard(motherboard MotherboardInfo) *HardwareInfoBuilder {
	b.hardware.Motherboard = motherboard
	return b
}

// Build finaliza a construção e retorna o objeto HardwareInfo
func (b *HardwareInfoBuilder) Build() HardwareInfo {
	return b.hardware
}
