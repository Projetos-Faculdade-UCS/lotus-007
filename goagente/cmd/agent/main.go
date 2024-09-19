package main

import (
	"encoding/json"
	"fmt"
	"goagente/internal/data/hardware"
	"goagente/internal/data/system"
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

	// Inicializa o builder de CoreInfoResult do pacote system
	builder := system.CoreInfoResultBuilder{}

	// Preenche automaticamente os dados de hostname e usuário
	_, err = builder.AutomaticPopulate() // Corrigido para não criar uma nova variável err
	if err != nil {
		log.Fatal("Erro ao preencher automaticamente CoreInfoResult:", err)
	}

	// Define o patrimônio manualmente
	builder.SetPatrimonio("12345")

	// Constrói o objeto CoreInfoResult
	coreInfo := builder.Build()

	// Converte o objeto CoreInfoResult para JSON
	coreInfoJSON, err := json.MarshalIndent(coreInfo, "", "    ")
	if err != nil {
		log.Fatal("Erro ao converter CoreInfoResult para JSON:", err)
	}

	// Exibe o JSON resultante
	fmt.Println("Informações de Core em JSON:")
	fmt.Println(string(coreInfoJSON))
}
