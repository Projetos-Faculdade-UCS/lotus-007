package security

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"goagente/internal/data/hardware"
	programs "goagente/internal/data/program"
	"goagente/internal/data/system"
)

// AddHMACToStruct adiciona o HMAC diretamente à struct, preservando a ordem dos campos
func AddHMACToStruct(data interface{}, secret string) (string, error) {
	// Serializa o objeto para JSON (sem o HMAC)
	jsonData, err := json.Marshal(data)
	if err != nil {
		return "", fmt.Errorf("erro ao serializar o objeto para JSON: %w", err)
	}

	// Gera o HMAC usando a chave secreta
	hmacHash := generateHMAC(jsonData, secret)

	// Adiciona o HMAC ao campo apropriado da struct
	switch v := data.(type) {
	case *hardware.HardwareInfo:
		v.HMAC = hmacHash
	case *system.CoreInfoResult:
		v.HMAC = hmacHash
	case *programs.ProgramInfo:
		v.HMAC = hmacHash
	default:
		return "", fmt.Errorf("tipo de dado não suportado para HMAC")
	}

	// Serializa novamente a struct com o HMAC no final
	finalJSON, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		return "", fmt.Errorf("erro ao serializar JSON final: %w", err)
	}

	return string(finalJSON), nil
}

// Função para gerar o HMAC
func generateHMAC(data []byte, secret string) string {
	h := hmac.New(sha256.New, []byte(secret))
	h.Write(data)
	return hex.EncodeToString(h.Sum(nil))
}
