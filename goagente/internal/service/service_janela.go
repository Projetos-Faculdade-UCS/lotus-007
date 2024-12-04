// Package service gerencia a execução e o ciclo de vida de serviços no sistema operacional Windows.
// Este arquivo contém a implementação do WindowsServiceManager, que executa tarefas periódicas enquanto o serviço está ativo.
package service

import (
	"log"
	"time"
)

// WindowsServiceManager é a implementação de ServiceManager para o sistema operacional Windows.
// Ele gerencia o ciclo de vida do serviço, incluindo inicialização, execução de tarefas e encerramento.
type WindowsServiceManager struct {
	config ServiceConfig // Configuração do serviço, incluindo nome e descrição
	exit   chan struct{} // Canal usado para sinalizar o encerramento do serviço
}

// Start inicia o serviço no sistema operacional Windows.
// Ele inicializa o canal `exit` e executa o método `run` em uma goroutine para processar tarefas em segundo plano.
//
// Retorna:
// - Um erro, caso a inicialização falhe (sempre retorna nil na implementação atual).
func (w *WindowsServiceManager) Start() error {
	w.exit = make(chan struct{})
	go w.run() // Executa a lógica do serviço em uma goroutine
	log.Printf("Serviço %s (Windows) iniciado.", w.config.Name)
	return nil
}

// Stop encerra o serviço no sistema operacional Windows.
// Ele fecha o canal `exit`, sinalizando para o método `run` interromper sua execução.
//
// Retorna:
// - Um erro, caso o encerramento falhe (sempre retorna nil na implementação atual).
func (w *WindowsServiceManager) Stop() error {
	close(w.exit) // Sinaliza para o método `run` encerrar
	log.Printf("Serviço %s (Windows) parado.", w.config.Name)
	return nil
}

// run executa a lógica principal do serviço em um loop.
// Ele utiliza um `time.Ticker` para realizar tarefas periódicas enquanto o serviço está ativo.
//
// Funcionalidade:
// - A cada 10 segundos, executa uma tarefa simulada e registra uma mensagem de log.
// - Interrompe sua execução quando o canal `exit` é fechado.
func (w *WindowsServiceManager) run() {
	ticker := time.NewTicker(10 * time.Second) // Define o intervalo para as tarefas
	defer ticker.Stop()                        // Garante que o ticker seja parado ao sair do loop

	for {
		select {
		case <-ticker.C:
			// Executa a tarefa periódica
			log.Printf("Serviço %s (Windows) executando tarefa...", w.config.Name)
		case <-w.exit:
			// Interrompe o loop ao receber sinal de encerramento
			log.Printf("Serviço %s (Windows) parando...", w.config.Name)
			return
		}
	}
}
