// Package main é o ponto de entrada do agente de orquestração.
// Ele inicializa e gerencia serviços, comunicação com APIs, e loops de orquestração para hardware e programas.
package main

import (
	"goagente/internal/communication" // Importa o pacote de comunicação
	"goagente/internal/logging"       // Importa o pacote de logging
	"goagente/internal/orchestration" // Importa o pacote de orquestração
	"goagente/internal/service"       // Importa o pacote de serviço
	"log"
	"time"
)

// main é a função principal que inicializa e executa o agente.
// Ela configura o sistema de logging, inicia serviços essenciais, e gerencia loops de orquestração.
func main() {
	// Inicializa o logger
	loggerFactory, err := logging.NewLoggerFactory()
	if err != nil {
		log.Fatal(err) // Finaliza o programa se ocorrer um erro ao iniciar o logger
	}
	defer loggerFactory.CloseLogger() // Garante que o logger será fechado ao final

	// Inicia o serviço em uma goroutine separada
	go func() {
		if err := service.RunService(); err != nil {
			log.Fatalf("Erro ao executar o serviço: %v", err)
		}
	}()

	// Inicializa o cliente de comunicação com a API
	client := communication.NewAPIClient("http://localhost:8000")
	poster := communication.NewInfoPoster(client) // Cria um InfoPoster para enviar informações à API

	// Define a chave secreta para operações seguras
	secretKey := "minhaChaveSecreta"

	// Inicializa os orquestradores para hardware e programas
	hardwareOrchestrator := orchestration.NewHardwareOrchestrator()
	programOrchestrator := orchestration.NewProgramOrchestrator()

	// Inicializa os mediadores usando os orquestradores e o InfoPoster
	hardwareMediator := orchestration.NewHardwareMediator(hardwareOrchestrator, poster, secretKey)
	coreMediator := orchestration.NewCoreMediator(poster, secretKey)
	programMediator := orchestration.NewProgramMediator(programOrchestrator, poster, secretKey)

	// Configura os loops de orquestração com seus respectivos mediadores e intervalos
	hardwareLoop := &orchestration.HardwareOrchestrationLoop{
		OrchestrationLoop: orchestration.OrchestrationLoop{
			Mediator: hardwareMediator,
			Interval: 10 * time.Second, // Intervalo de 10 segundos
		},
	}

	coreLoop := &orchestration.CoreOrchestrationLoop{
		OrchestrationLoop: orchestration.OrchestrationLoop{
			Mediator: coreMediator,
			Interval: 10 * time.Second, // Intervalo de 10 segundos
		},
	}

	programLoop := &orchestration.ProgramOrchestrationLoop{
		OrchestrationLoop: orchestration.OrchestrationLoop{
			Mediator: programMediator,
			Interval: 10 * time.Second, // Intervalo de 10 segundos
		},
	}

	// Inicia o loop de orquestração do core
	go coreLoop.Start()

	// Aguarda 10 segundos antes de iniciar o loop de hardware
	time.Sleep(10 * time.Second)
	go hardwareLoop.Start()

	// Inicia o loop de orquestração de programas
	go programLoop.Start()

	// Mantém o programa em execução indefinidamente
	select {}
}
