package utils

import (
	"encoding/json"
	"fmt"
)

func BytesToGigabytes(bytes uint64) float64 {
	const bytesInGigabyte = 1024 * 1024 * 1024
	return float64(bytes) / bytesInGigabyte
}

// SerializeToJSON converte um objeto para uma string JSON
func SerializeToJSON(data interface{}) (string, error) {
	jsonData, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		return "", fmt.Errorf("erro ao converter objeto para JSON: %w", err)
	}
	return string(jsonData), nil
}
