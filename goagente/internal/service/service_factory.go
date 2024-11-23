// Package service gerencia a instalação, execução e desinstalação de serviços para diferentes sistemas operacionais.
// Ele fornece suporte para configurar serviços como Windows Services ou Linux Daemons.
package service

import (
	"fmt"
	"log"
	"os"
	"runtime"

	"github.com/kardianos/service"
)

// ServiceConfig define as configurações do serviço, incluindo nome, descrição e comportamento de inicialização.
type ServiceConfig struct {
	Name        string // Nome interno do serviço
	DisplayName string // Nome exibido do serviço
	Description string // Descrição do serviço
	AutoStart   bool   // Indica se o serviço deve ser iniciado automaticamente
}

// NewDefaultServiceConfig cria uma configuração com valores padrão para o serviço.
//
// Retorna:
// - Uma instância de ServiceConfig com valores pré-definidos.
func NewDefaultServiceConfig() ServiceConfig {
	return ServiceConfig{
		Name:        "LotusAgent",
		DisplayName: "Lotus",
		Description: "Agente",
		AutoStart:   true,
	}
}

// ServiceManager define uma interface para gerenciar o ciclo de vida de um serviço.
// Implementações específicas para diferentes sistemas operacionais devem implementar esta interface.
type ServiceManager interface {
	// Start inicia o serviço.
	Start() error

	// Stop encerra o serviço.
	Stop() error
}

// NewServiceManager é uma factory que retorna a implementação apropriada de ServiceManager com base no sistema operacional.
//
// Parâmetros:
// - config: Configuração do serviço.
//
// Retorna:
// - Uma implementação de ServiceManager compatível com o sistema operacional.
// - Um erro, caso o sistema operacional não seja suportado.
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

// program é uma estrutura que gerencia o ciclo de vida do serviço.
// Ele é usado pela biblioteca kardianos/service para iniciar e parar o serviço.
type program struct {
	manager ServiceManager // Gerenciador responsável por iniciar e parar o serviço
}

// Start inicializa o serviço.
//
// Parâmetros:
// - s: Instância do serviço fornecida pela biblioteca kardianos/service.
//
// Retorna:
// - Um erro, caso a inicialização do serviço falhe.
func (p *program) Start(s service.Service) error {
	return p.manager.Start()
}

// Stop encerra o serviço.
//
// Parâmetros:
// - s: Instância do serviço fornecida pela biblioteca kardianos/service.
//
// Retorna:
// - Um erro, caso a interrupção do serviço falhe.
func (p *program) Stop(s service.Service) error {
	return p.manager.Stop()
}

// RunService gerencia a instalação, execução e desinstalação do serviço.
//
// Funcionalidade:
// - Se `install` for passado como argumento de linha de comando, instala o serviço.
// - Se `uninstall` for passado como argumento de linha de comando, desinstala o serviço.
// - Caso contrário, executa o serviço.
//
// Retorna:
// - Um erro, caso qualquer operação falhe.
func RunService() error {
	// Cria a configuração padrão do serviço
	config := NewDefaultServiceConfig()

	// Configura os parâmetros do serviço para a biblioteca kardianos/service
	svcConfig := &service.Config{
		Name:        config.Name,
		DisplayName: config.DisplayName,
		Description: config.Description,
	}

	// Cria o gerenciador de serviço apropriado para o sistema operacional
	manager, err := NewServiceManager(config)
	if err != nil {
		return fmt.Errorf("erro ao criar o gerenciador do serviço: %v", err)
	}

	// Cria a estrutura do programa e associa ao serviço
	prg := &program{manager: manager}
	s, err := service.New(prg, svcConfig)
	if err != nil {
		return fmt.Errorf("erro ao criar o serviço: %v", err)
	}

	// Gerencia a instalação ou desinstalação do serviço com base nos argumentos da linha de comando
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

	// Executa o serviço
	return s.Run()
}
