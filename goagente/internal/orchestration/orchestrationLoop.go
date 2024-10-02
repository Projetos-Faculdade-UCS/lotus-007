package orchestration

import (
	"fmt"
	"goagente/internal/logging"
	"time"
)

type OrchestrationLoop struct {
	Mediator   MediatorInterface // Interface para o Mediator específico
	Patrimonio string
	Interval   time.Duration // Intervalo de tempo para o loop
}

type HardwareOrchestrationLoop struct {
	OrchestrationLoop
}

type CoreOrchestrationLoop struct {
	OrchestrationLoop
}

type ProgramOrchestrationLoop struct {
	OrchestrationLoop
}

type MediatorInterface interface {
	OrchestrateAndPost(patrimonio string) error
}

// Start inicia o loop de orquestração
func (o *OrchestrationLoop) Start() {
	for {
		err := o.Mediator.OrchestrateAndPost(o.Patrimonio)
		if err != nil {
			logging.Error(fmt.Errorf("erro ao orquestrar e enviar dados: %w", err))
		}
		time.Sleep(o.Interval)
	}
}
