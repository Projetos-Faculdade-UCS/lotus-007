package service

import (
	"log"
	"time"
)

type WindowsServiceManager struct {
	config ServiceConfig
	exit   chan struct{}
}

func (w *WindowsServiceManager) Start() error {
	w.exit = make(chan struct{})
	go w.run()
	log.Printf("Serviço %s (Windows) iniciado.", w.config.Name)
	return nil
}

func (w *WindowsServiceManager) Stop() error {
	close(w.exit)
	log.Printf("Serviço %s (Windows) parado.", w.config.Name)
	return nil
}

func (w *WindowsServiceManager) run() {
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop() // Garante que o ticker seja parado ao sair

	for {
		select {
		case <-ticker.C:
			log.Printf("Serviço %s (Windows) executando tarefa...", w.config.Name)
		case <-w.exit:
			log.Printf("Serviço %s (Windows) parando...", w.config.Name)
			return
		}
	}
}
