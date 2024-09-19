package main

import (
	"encoding/json"
	"fmt"
	"goagente/internal/data/hardware"
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

	// Inicializa os retrievers de hardware (RAM, Discos, Processadores e Placa-mãe)
	ramRetriever := hardware.WindowsRAMRetriever{}
	diskRetriever := hardware.WindowsDiskRetriever{}
	processorRetriever := hardware.WindowsProcessorRetriever{}
	motherboardRetriever := hardware.WindowsMotherboardRetriever{}

	// Inicializa o orquestrador com os retrievers
	orchestrator := orchestration.NewHardwareOrchestrator(ramRetriever, diskRetriever, processorRetriever, motherboardRetriever)

	// Executa o método Orchestrate para montar o HardwareInfo
	patrimonio := "12345" // Exemplo de valor de patrimônio
	hardwareInfo, err := orchestrator.Orchestrate(patrimonio)
	if err != nil {
		log.Fatal("Erro ao orquestrar informações de hardware:", err)
	}

	// Converte o objeto HardwareInfo em JSON
	hardwareInfoJSON, err := json.MarshalIndent(hardwareInfo, "", "    ")
	if err != nil {
		log.Fatal("Erro ao converter HardwareInfo para JSON:", err)
	}

	// Exibe o JSON resultante
	fmt.Println("Informações de Hardware em JSON:")
	fmt.Println(string(hardwareInfoJSON))
}
