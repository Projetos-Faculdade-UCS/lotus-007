package main

import (
	"encoding/json"
	"fmt"
	communication "goagente/internal/communication"
	"goagente/internal/logging"
	"goagente/internal/orchestration"
	"log"
)

func main() {
	// Inicializa o logger
	loggerFactory, err := logging.NewLoggerFactory()
	if err != nil {
		log.Fatal(err)
	}
	defer loggerFactory.CloseLogger()

	client := communication.NewAPIClient("https://api.meuservidor.com")
	poster := communication.NewInfoPoster(client)

	logging.Info("Iniciando a execução do HardwareInfoBuilder")

	orchestrator := orchestration.NewHardwareOrchestrator()

	// Executa o método Orchestrate para montar o HardwareInfo
	patrimonio := "12345"
	hardwareInfo, err := orchestrator.Orchestrate(patrimonio)
	if err != nil {
		log.Fatal("Erro ao orquestrar informações de hardware:", err)
	}

	// Converte o objeto HardwareInfo para JSON e exibe
	hardwareInfoJSON, err := json.MarshalIndent(hardwareInfo, "", "    ")
	if err != nil {
		log.Fatal("Erro ao converter HardwareInfo para JSON:", err)
	}

	err = poster.PostHardwareInfo("hardware", string(hardwareInfoJSON))
	if err != nil {
		fmt.Println("Erro ao enviar informações de core:", err)
	}

	// Orquestra as informações de CoreInfoResult
	coreInfo, err := orchestration.OrchestrateCoreInfo(patrimonio)
	if err != nil {
		log.Fatal("Erro ao orquestrar informações de CoreInfoResult:", err)
	}

	// Converte o objeto CoreInfoResult para JSON
	coreInfoJSON, err := json.MarshalIndent(coreInfo, "", "    ")
	if err != nil {
		log.Fatal("Erro ao converter CoreInfoResult para JSON:", err)
	}

	err = poster.PostCoreInfo("core", string(coreInfoJSON))
	if err != nil {
		fmt.Println("Erro ao enviar informações de core:", err)
	}

	orchestratorp := orchestration.NewProgramOrchestrator()

	// Executa o método Orchestrate para montar o ProgramInfo
	programInfo, err := orchestratorp.Orchestrate(patrimonio)
	if err != nil {
		log.Fatal("Erro ao orquestrar informações de programas instalados:", err)
	}

	// Converte o objeto ProgramInfo para JSON e exibe
	programInfoJSON, err := json.MarshalIndent(programInfo, "", "    ")
	if err != nil {
		log.Fatal("Erro ao converter ProgramInfo para JSON:", err)
	}

	err = poster.PostProgramInfo("program", string(programInfoJSON))
	if err != nil {
		fmt.Println("Erro ao enviar informações de programas:", err)
	}

}
