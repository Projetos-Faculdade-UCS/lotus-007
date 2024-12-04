// Package programs fornece funcionalidades para coletar informações sobre os programas instalados no sistema.
// Este arquivo contém a implementação de WindowsProgramsRetriever para coletar programas no sistema operacional Windows.
package programs

import (
	"golang.org/x/sys/windows/registry"
)

// WindowsProgramsRetriever é a implementação de ProgramRetriever para o sistema operacional Windows.
// Ele utiliza o registro do Windows para coletar informações sobre programas instalados.
type WindowsProgramsRetriever struct{}

// GetInstalledPrograms coleta informações sobre os programas instalados no Windows.
// Ele verifica chaves específicas no registro do Windows para obter a lista de programas.
//
// Retorna:
// - Um slice de Program contendo as informações dos programas instalados.
// - Um erro, caso ocorra algum problema durante a leitura do registro.
func (p WindowsProgramsRetriever) GetInstalledPrograms() ([]Program, error) {
	var programs []Program

	// Caminhos do registro onde os programas instalados estão listados
	keys := []string{
		`SOFTWARE\Microsoft\Windows\CurrentVersion\Uninstall`,
		`SOFTWARE\WOW6432Node\Microsoft\Windows\CurrentVersion\Uninstall`,
	}

	for _, key := range keys {
		// Abre a chave do registro para leitura
		k, err := registry.OpenKey(registry.LOCAL_MACHINE, key, registry.READ)
		if err != nil {
			continue // Ignora erros ao abrir a chave
		}
		defer k.Close()

		// Lê os nomes das subchaves
		names, err := k.ReadSubKeyNames(-1)
		if err != nil {
			continue // Ignora erros ao ler subchaves
		}

		for _, name := range names {
			// Abre a subchave para leitura
			subKey, err := registry.OpenKey(k, name, registry.READ)
			if err != nil {
				continue // Ignora erros ao abrir a subchave
			}
			defer subKey.Close()

			// Lê o valor "DisplayName" da subchave
			displayName, _, err := subKey.GetStringValue("DisplayName")
			if err != nil {
				continue // Ignora subchaves sem "DisplayName"
			}

			// Lê o valor "DisplayVersion" da subchave
			displayVersion, _, err := subKey.GetStringValue("DisplayVersion")
			if err != nil {
				displayVersion = "N/A" // Define "N/A" caso "DisplayVersion" não exista
			}

			// Adiciona o programa à lista
			programs = append(programs, Program{Name: displayName, Version: displayVersion})
		}
	}

	return programs, nil
}
