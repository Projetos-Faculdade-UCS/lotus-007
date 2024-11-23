// Package orchestration gerencia loops de orquestração que executam tarefas periódicas de coleta e envio de dados.
// Este arquivo contém a implementação de loops genéricos e específicos para hardware, core e programas.
package orchestration

import (
	"fmt"
	"goagente/internal/logging"
	"time"
)

// OrchestrationLoop define um loop genérico para executar tarefas de orquestração periódicas.
// Ele utiliza um Mediator para realizar a coleta e o envio de dados.
type OrchestrationLoop struct {
	Mediator   MediatorInterface // Interface para o Mediator específico
	Patrimonio string            // Identificação do patrimônio para orquestração
	Interval   time.Duration     // Intervalo de tempo entre as execuções do loop
}

// HardwareOrchestrationLoop representa um loop específico para a orquestração de hardware.
type HardwareOrchestrationLoop struct {
	OrchestrationLoop
}

// CoreOrchestrationLoop representa um loop específico para a orquestração de informações de Core.
type CoreOrchestrationLoop struct {
	OrchestrationLoop
}

// ProgramOrchestrationLoop representa um loop específico para a orquestração de programas.
type ProgramOrchestrationLoop struct {
	OrchestrationLoop
}

// MediatorInterface define a interface para Mediators usados nos loops de orquestração.
// Cada Mediator deve implementar o método OrchestrateAndPost.
type MediatorInterface interface {
	// OrchestrateAndPost realiza a coleta de informações e o envio para o servidor.
	//
	// Parâmetros:
	// - patrimonio: Identificação do patrimônio para ser usada na orquestração.
	//
	// Retorna:
	// - Um erro, caso ocorra algum problema durante a execução.
	OrchestrateAndPost(patrimonio string) error
}

// Start inicia o loop de orquestração, executando o Mediator periodicamente.
// O loop continuará executando indefinidamente.
//
// Funcionalidade:
// - A cada intervalo definido, chama o método OrchestrateAndPost do Mediator.
// - Registra erros no sistema de logging, caso ocorram.
//
// Observação:
// - Esta função bloqueia a execução indefinidamente.
// - Recomenda-se executá-la em uma goroutine separada.
//
// Exemplo de Uso:
//
//	go loop.Start()
func (o *OrchestrationLoop) Start() {
	for {
		err := o.Mediator.OrchestrateAndPost(o.Patrimonio)
		if err != nil {
			logging.Error(fmt.Errorf("erro ao orquestrar e enviar dados: %w", err))
		}
		time.Sleep(o.Interval) // Aguarda o intervalo antes da próxima execução
	}
}
