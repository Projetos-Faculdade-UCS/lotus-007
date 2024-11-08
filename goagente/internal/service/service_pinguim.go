package service

import (
	"log"
	"time"
)

type LinuxServiceManager struct {
	config ServiceConfig
	exit   chan struct{}
}

func (l *LinuxServiceManager) Start() error {
	l.exit = make(chan struct{})
	go l.run()
	log.Printf("Serviço %s (Linux) iniciado.", l.config.Name)
	return nil
}

func (l *LinuxServiceManager) Stop() error {
	close(l.exit)
	log.Printf("Serviço %s (Linux) parado.", l.config.Name)
	return nil
}

func (l *LinuxServiceManager) run() {
	ticker := time.NewTicker(10 * time.Second)
	for {
		select {
		case <-ticker.C:
			log.Printf("Serviço %s (Linux) executando tarefa...", l.config.Name)
		case <-l.exit:
			ticker.Stop()
			return
		}
	}
}
