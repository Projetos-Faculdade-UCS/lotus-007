package service

import (
	"fmt"
	"log"
	"os"
	"runtime"

	"github.com/kardianos/service"
)

// ServiceConfig para definir configurações do serviço com valores padrão
type ServiceConfig struct {
	Name        string
	DisplayName string
	Description string
	AutoStart   bool
}

// NewDefaultServiceConfig cria uma configuração com valores padrão
func NewDefaultServiceConfig() ServiceConfig {
	return ServiceConfig{
		Name:        "LotusAgent",
		DisplayName: "Lotus",
		Description: "Agente",
		AutoStart:   true,
	}
}

// ServiceManager interface com métodos Start e Stop
type ServiceManager interface {
	Start() error
	Stop() error
}

// NewServiceManager factory que retorna a implementação correta para o SO
func NewServiceManager(config ServiceConfig) (ServiceManager, error) {
	so := runtime.GOOS

	switch so {
	case "windows":
		return &WindowsServiceManager{config: config}, nil
	case "linux":
		return &LinuxServiceManager{config: config}, nil
	default:
		return nil, fmt.Errorf("sistema operacional não suportado: %s", so)
	}
}

// Estrutura program para gerenciar o ciclo de vida do serviço
type program struct {
	manager ServiceManager
}

// Start inicializa o serviço
func (p *program) Start(s service.Service) error {
	// Inicia o gerenciador de serviço
	return p.manager.Start()
}

// Stop encerra o serviço
func (p *program) Stop(s service.Service) error {
	// Encerra o gerenciador de serviço
	return p.manager.Stop()
}

// RunService gerencia a instalação, execução e desinstalação do serviço
func RunService() error {
	config := NewDefaultServiceConfig()

	svcConfig := &service.Config{
		Name:        config.Name,
		DisplayName: config.DisplayName,
		Description: config.Description,
	}

	manager, err := NewServiceManager(config)
	if err != nil {
		return fmt.Errorf("erro ao criar o gerenciador do serviço: %v", err)
	}

	prg := &program{manager: manager}
	s, err := service.New(prg, svcConfig)
	if err != nil {
		return fmt.Errorf("erro ao criar o serviço: %v", err)
	}

	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "install":
			err := s.Install()
			if err != nil {
				return fmt.Errorf("falha ao instalar o serviço: %v", err)
			}
			log.Println("Serviço instalado com sucesso.")
			return nil
		case "uninstall":
			err := s.Uninstall()
			if err != nil {
				return fmt.Errorf("falha ao desinstalar o serviço: %v", err)
			}
			log.Println("Serviço desinstalado com sucesso.")
			return nil
		}
	}

	return s.Run()
}
