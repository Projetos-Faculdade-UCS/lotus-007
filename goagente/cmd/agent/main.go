package main

import (
	"goagente/internal/communication" // Importa o pacote do orquestrador de hardware
	// Importa o pacote do orquestrador de programas
	"goagente/internal/logging"
	"goagente/internal/orchestration"
	"log"
	"time"
)

func main() {
	// Inicializa o logger
	loggerFactory, err := logging.NewLoggerFactory()
	if err != nil {
		log.Fatal(err)
	}
	defer loggerFactory.CloseLogger()

	// Inicializa a camada de comunicação
	client := communication.NewAPIClient("https://api.meuservidor.com")
	poster := communication.NewInfoPoster(client)

	// Define o patrimônio e a chave secreta para HMAC
	patrimonio := "12345"
	secretKey := "minhaChaveSecreta"

	// Inicializa o orquestrador de hardware e programas
	hardwareOrchestrator := orchestration.NewHardwareOrchestrator()
	programOrchestrator := orchestration.NewProgramOrchestrator()

	// Inicializa os mediadores com a chave secreta
	hardwareMediator := orchestration.NewHardwareMediator(hardwareOrchestrator, poster, secretKey)
	coreMediator := orchestration.NewCoreMediator(poster, secretKey)
	programMediator := orchestration.NewProgramMediator(programOrchestrator, poster, secretKey)

	// Cria os loops de orquestração com seus respectivos mediadores e intervalos
	hardwareLoop := &orchestration.HardwareOrchestrationLoop{
		OrchestrationLoop: orchestration.OrchestrationLoop{
			Mediator:   hardwareMediator,
			Patrimonio: patrimonio,
			Interval:   60 * time.Second, // Define o intervalo de 5 segundos
		},
	}

	coreLoop := &orchestration.CoreOrchestrationLoop{
		OrchestrationLoop: orchestration.OrchestrationLoop{
			Mediator:   coreMediator,
			Patrimonio: patrimonio,
			Interval:   60 * time.Second, // Define o intervalo de 5 segundos
		},
	}

	programLoop := &orchestration.ProgramOrchestrationLoop{
		OrchestrationLoop: orchestration.OrchestrationLoop{
			Mediator:   programMediator,
			Patrimonio: patrimonio,
			Interval:   60 * time.Second, // Define o intervalo de 5 segundos
		},
	}

	// Inicia os loops de orquestração em goroutines
	go hardwareLoop.Start()
	go coreLoop.Start()
	go programLoop.Start()

	// Mantém o programa em execução
	select {}
}
