package main

import (
	"fmt"
	"goagente/internal/communication"
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

	// Inicializa os mediadores com a chave secreta
	hardwareMediator := orchestration.NewHardwareMediator(orchestration.NewHardwareOrchestrator(), poster, secretKey)
	coreMediator := orchestration.NewCoreMediator(poster, secretKey)
	programMediator := orchestration.NewProgramMediator(orchestration.NewProgramOrchestrator(), poster, secretKey)

	// Executa as funções de orquestração em goroutines com looping de 5 segundos
	go loopHardwareOrchestration(hardwareMediator, patrimonio)
	go loopCoreOrchestration(coreMediator, patrimonio)
	go loopProgramOrchestration(programMediator, patrimonio)

	// Mantém o programa em execução
	select {}
}

// loopHardwareOrchestration faz a orquestração de hardware a cada 5 segundos
func loopHardwareOrchestration(hardwareMediator *orchestration.HardwareMediator, patrimonio string) {
	for {
		err := hardwareMediator.OrchestrateAndPost(patrimonio)
		if err != nil {
			newErr := fmt.Errorf("erro ao orquestrar e enviar informações de hardware: %w", err)
			logging.Error(newErr)
		}
		time.Sleep(5 * time.Second) // Pausa de 5 segundos
	}
}

// loopCoreOrchestration faz a orquestração de core a cada 5 segundos
func loopCoreOrchestration(coreMediator *orchestration.CoreMediator, patrimonio string) {
	for {
		err := coreMediator.OrchestrateAndPost(patrimonio)
		if err != nil {
			newErr := fmt.Errorf("erro ao orquestrar e enviar informações de core: %w", err)
			logging.Error(newErr)
		}
		time.Sleep(5 * time.Second) // Pausa de 5 segundos
	}
}

// loopProgramOrchestration faz a orquestração de programas a cada 5 segundos
func loopProgramOrchestration(programMediator *orchestration.ProgramMediator, patrimonio string) {
	for {
		err := programMediator.OrchestrateAndPost(patrimonio)
		if err != nil {
			newErr := fmt.Errorf("erro ao orquestrar e enviar informações de programas: %w", err)
			logging.Error(newErr)
		}
		time.Sleep(5 * time.Second) // Pausa de 5 segundos
	}
}
