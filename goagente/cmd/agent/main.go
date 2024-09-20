package main

import (
	"encoding/json"
	"fmt"
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
	fmt.Println("Informações de Hardware em JSON:")
	fmt.Println(string(hardwareInfoJSON))
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

	// Exibe o JSON resultante
	fmt.Println("Informações de Core em JSON:")
	fmt.Println(string(coreInfoJSON))
}
