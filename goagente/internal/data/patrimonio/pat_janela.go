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

type WindowsPatRetriever struct{}

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

// ExtrairSequenciaFinal extrai a sequência numérica final ou gera um número negativo aleatório de 6 dígitos
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

// fileExists verifica se o arquivo já existe
func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	return err == nil && !info.IsDir()
}

// readSequenceFromFile lê a sequência do arquivo pat.txt
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

// writeSequenceToFile escreve a sequência no arquivo pat.txt
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
