package main

import (
	"fmt"
	"goagente/internal/logging"
	"log"
)

func main() {

	loggerFactory, err := logging.NewLoggerFactory()
	if err != nil {
		log.Fatal(err)
	}

	// Certifique-se de fechar os loggers no final
	defer loggerFactory.CloseLogger()

	logging.Info("Executando DoWork em mypackage")
	logging.Debug("Informações detalhadas sobre DoWork")
	logging.Error(fmt.Errorf("erro durante a execução de DoWork"))

	select {}
}
