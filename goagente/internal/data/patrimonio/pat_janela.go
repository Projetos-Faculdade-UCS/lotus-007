// Package patrimonio fornece funcionalidades para coletar e gerenciar informações sobre o patrimônio do sistema.
// Este arquivo contém a implementação de WindowsPatRetriever, que coleta o patrimônio com base no hostname ou arquivo local.
package patrimonio

import (
	"bufio"
	"fmt"
	"goagente/internal/logging"
	"math/rand"
	"os"
	"regexp"
	"strings"
	"time"
)

// WindowsPatRetriever é a implementação de PatRetriever para o sistema operacional Windows.
// Ele obtém o patrimônio com base no hostname do computador ou em um arquivo local.
type WindowsPatRetriever struct{}

// GetCurrentPat retorna o patrimônio atual no Windows.
// O patrimônio é extraído do hostname ou de um arquivo local (pat.txt).
// Caso o arquivo não exista, ele será criado com a sequência extraída do hostname.
//
// Retorna:
// - Uma string representando o patrimônio do sistema.
// - Um erro, caso a operação falhe.
func (WindowsPatRetriever) GetCurrentPat() (string, error) {
	// Verificar se o arquivo pat.txt existe e contém uma sequência
	patFile := "pat.txt"
	if fileExists(patFile) {
		sequence, err := readSequenceFromFile(patFile)
		if err == nil && sequence != "" {
			logging.Info("Sequência encontrada no arquivo pat.txt.")
			return sequence, nil
		}
	}

	// Obter o hostname do computador
	hostname, err := os.Hostname()
	if err != nil {
		newErr := fmt.Errorf("erro ao obter o hostname no Windows: %v", err)
		logging.Error(newErr)
		return "", err
	}

	// Extrair a sequência numérica final e salvar no arquivo pat.txt
	sequence := ExtrairSequenciaFinal(hostname)
	err = writeSequenceToFile(patFile, sequence)
	if err != nil {
		logging.Error(fmt.Errorf("erro ao escrever a sequência no arquivo pat.txt: %v", err))
	}
	return sequence, nil
}

// ExtrairSequenciaFinal extrai a sequência numérica final do nome do computador ou gera um número negativo aleatório de 6 dígitos.
//
// Parâmetros:
// - nomeComputador: Nome do computador.
//
// Retorna:
// - Uma string contendo a sequência extraída ou um número negativo gerado aleatoriamente.
func ExtrairSequenciaFinal(nomeComputador string) string {
	re := regexp.MustCompile(`[0-9]+$`)
	sequence := re.FindString(nomeComputador)

	if sequence == "" {
		logging.Info("Nenhuma sequência numérica encontrada no nome do computador.")
		rng := rand.New(rand.NewSource(time.Now().UnixNano()))
		return fmt.Sprintf("-%06d", rng.Intn(900000)+100000)
	}
	return sequence
}

// fileExists verifica se o arquivo especificado existe.
//
// Parâmetros:
// - filename: Nome do arquivo a ser verificado.
//
// Retorna:
// - True, se o arquivo existir.
// - False, caso contrário.
func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	return err == nil && !info.IsDir()
}

// readSequenceFromFile lê a sequência de um arquivo especificado.
//
// Parâmetros:
// - filename: Nome do arquivo de onde a sequência será lida.
//
// Retorna:
// - Uma string contendo a sequência lida.
// - Um erro, caso a leitura falhe.
func readSequenceFromFile(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		logging.Error(fmt.Errorf("erro ao ler o arquivo pat.txt: %v", err))
		return "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	if scanner.Scan() {
		return strings.TrimSpace(scanner.Text()), nil
	}
	return "", scanner.Err()
}

// writeSequenceToFile escreve a sequência especificada em um arquivo.
//
// Parâmetros:
// - filename: Nome do arquivo onde a sequência será escrita.
// - sequence: Sequência a ser escrita no arquivo.
//
// Retorna:
// - Um erro, caso a escrita falhe.
func writeSequenceToFile(filename, sequence string) error {
	file, err := os.Create(filename)
	if err != nil {
		logging.Error(fmt.Errorf("erro ao escrever o arquivo pat.txt: %v", err))
		return err
	}
	defer file.Close()

	_, err = file.WriteString(sequence)
	return err
}
