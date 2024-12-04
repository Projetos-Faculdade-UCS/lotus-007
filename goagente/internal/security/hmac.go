// Package security fornece funcionalidades para adicionar segurança a dados utilizando HMAC.
// Este arquivo contém a implementação da função AddHMACToStruct para adicionar HMAC a structs.
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

// AddHMACToStruct adiciona um HMAC a uma struct específica, preservando a ordem dos campos.
//
// Funcionalidade:
// - Serializa a struct em JSON, gerando um HMAC com base nos dados serializados e uma chave secreta.
// - Adiciona o HMAC ao campo apropriado da struct.
// - Serializa novamente a struct com o HMAC incluído.
//
// Parâmetros:
// - data: Ponteiro para a struct que receberá o HMAC. Tipos suportados:
//   - hardware.HardwareInfo
//   - system.CoreInfoResult
//   - programs.ProgramInfo
//
// - secret: Chave secreta utilizada para gerar o HMAC.
//
// Retorna:
// - Uma string contendo a representação JSON da struct com o HMAC.
// - Um erro, caso ocorra algum problema na serialização ou se o tipo da struct não for suportado.
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

// generateHMAC gera um HMAC para os dados fornecidos usando a chave secreta.
//
// Parâmetros:
// - data: Dados em formato de byte para os quais o HMAC será gerado.
// - secret: Chave secreta utilizada para gerar o HMAC.
//
// Retorna:
// - Uma string representando o HMAC gerado em formato hexadecimal.
func generateHMAC(data []byte, secret string) string {
	h := hmac.New(sha256.New, []byte(secret))
	h.Write(data)
	return hex.EncodeToString(h.Sum(nil))
}
